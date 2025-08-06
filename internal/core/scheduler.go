package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"hegelscheduler/internal/data"
	"hegelscheduler/internal/dto"
	"hegelscheduler/internal/model"
	"hegelscheduler/internal/queue"
	"log"
	"time"
)

type HegelScheduler struct {
	IsLeader         bool
	jobChan          chan *model.Job
	stopChan         chan bool
	scheduler        gocron.Scheduler
	productor        queue.Productor
	jobRepo          data.JobRepo
	jobExecutionRepo data.JobExecutionRepo
	jobMap           map[uint64]*gocron.Job
	checkPointTime   *time.Time
}

func NewHegelScheduler(productor queue.Productor, jobRepo data.JobRepo, jobExecutionRepo data.JobExecutionRepo) *HegelScheduler {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	return &HegelScheduler{
		jobChan:          make(chan *model.Job, 100),
		stopChan:         make(chan bool),
		IsLeader:         false,
		scheduler:        scheduler,
		productor:        productor,
		jobRepo:          jobRepo,
		jobExecutionRepo: jobExecutionRepo,
		jobMap:           make(map[uint64]*gocron.Job),
	}
}

type JobType string

const (
	simple     JobType = "simple"
	cron       JobType = "cron"
	fixedDelay JobType = "fixedDelay"
)

// Elect Elect per 5 second
func (s *HegelScheduler) Elect() {
	for {
		select {
		case <-time.After(time.Second * 5):
			// todo: 选举
			s.Start()
		case <-s.stopChan:
			return
		}
	}
}

// Start the HegelScheduler
func (s *HegelScheduler) Start() {
	s.scheduler.Start()
	go s.Consumer()
	go func() {
		if err := s.Poll(); err != nil {
			log.Println(err.Error())
		}
	}()
}

// Consumer New Jobs
func (s *HegelScheduler) Consumer() {
	for {
		select {
		case <-s.stopChan:
			return
		case job := <-s.jobChan:
			func(job *model.Job) {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Recovered in HegelScheduler", r)
					}
					if err := s.AddJob(job); err != nil {
						log.Println(err)
					}
				}()
			}(job)
		}
	}
}

func (s *HegelScheduler) AddJob(job *model.Job) error {
	var (
		goJob gocron.Job
		err   error
	)
	if _, exist := s.jobMap[job.ID]; exist {
		// todo：暂时不处理job更新情况
		return errors.New("job exist")
	}
	switch JobType(job.Type) {
	case simple:
		goJob, err = s.scheduler.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(*job.RunAt)), gocron.NewTask(func(job *model.Job) {
			s.Execute(job)
		}, job), gocron.WithName(job.Name))
	case cron:
		goJob, err = s.scheduler.NewJob(gocron.CronJob(*job.CronExpr, true), gocron.NewTask(func(job *model.Job) {
			s.Execute(job)
		}, job), gocron.WithName(job.Name))
	}
	if err != nil {
		return err
	}
	if goJob == nil {
		return errors.New("create new job failed")
	} else {
		s.jobMap[job.ID] = &goJob
	}
	return nil
}

// Execute jobExecution event to queue
func (s *HegelScheduler) Execute(job *model.Job) error {
	jobExection := model.JobExecution{
		JobID:         job.ID,
		ScheduledTime: time.Now(),
		StartTime:     nil,
		EndTime:       nil,
		Status:        model.JobExecutionStatusReady,
		Result:        "",
		WorkerID:      "",
		WorkerIP:      "",
		CreatedAt:     time.Now(),
	}
	if err := s.jobExecutionRepo.Create(&jobExection); err != nil {
		return err
	}
	dto := dto.JobExectionDto{
		JobExectionId: jobExection.ID,
		JobId:         job.ID,
		Name:          job.Name,
		Description:   job.Description,
		RetryCount:    job.RetryCount,
		RetryInterval: job.RetryInterval,
		Timeout:       job.Timeout,
		Payload:       job.Payload,
		TargetURL:     job.TargetURL,
		Method:        job.Method,
		Headers:       job.Headers,
	}
	if err := s.productor.Publish(dto, Queue); err != nil {
		return err
	}
	return nil
}

// Stop the scheduler
func (s *HegelScheduler) Stop() error {
	close(s.stopChan)
	return nil
}

// Poll new jobs per second and push it into scheduler chan
func (s *HegelScheduler) Poll() error {
	for {
		select {
		case <-time.After(time.Second):
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Poller recovered:", r)
					}
				}()
				newJobs, err := s.GetAvailableJobs()
				if err != nil {
					fmt.Println(err)
				}
				for _, job := range newJobs {
					s.jobChan <- job
				}
			}()
		case <-s.stopChan:
			return nil
		}
	}
}

// GetAvailableJobs get jobs from db
func (s *HegelScheduler) GetAvailableJobs() ([]*model.Job, error) {
	var (
		ctx  = context.Background()
		err  error
		jobs []*model.Job
	)
	if s.checkPointTime == nil {
		jobs, err = s.jobRepo.GetAvailableJobs(ctx)
	} else {
		jobs, err = s.jobRepo.GetNewAvailableJobs(ctx, *s.checkPointTime)
	}
	if err != nil {
		return nil, err
	}
	if latestJob, exist := findLatest(jobs, func(job *model.Job) time.Time {
		return job.UpdatedAt
	}); exist {
		s.checkPointTime = &latestJob.UpdatedAt
	}
	return jobs, nil
}

func findLatest[T any](items []T, getTime func(T) time.Time) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}

	latest := items[0]
	maxTime := getTime(latest)

	for _, item := range items[1:] {
		t := getTime(item)
		if t.After(maxTime) {
			latest = item
			maxTime = t
		}
	}

	return latest, true
}
