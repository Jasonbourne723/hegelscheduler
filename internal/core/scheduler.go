package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"hegelscheduler/internal/data"
	"hegelscheduler/internal/model"
	"hegelscheduler/internal/queue"
	"log"
	"time"
)

type HegelScheduler struct {
	IsLeader  bool
	jobChan   chan *model.Job
	stopChan  chan bool
	scheduler gocron.Scheduler
	productor queue.Productor
	jobRepo   data.JobRepo
}

func NewHegelScheduler(productor queue.Productor, jobRepo data.JobRepo) *HegelScheduler {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	return &HegelScheduler{
		jobChan:   make(chan *model.Job, 100),
		stopChan:  make(chan bool),
		IsLeader:  false,
		scheduler: scheduler,
		productor: productor,
		jobRepo:   jobRepo,
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
	switch JobType(job.Type) {
	case simple:
		goJob, err = s.scheduler.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(*job.RunAt)), gocron.NewTask(func(job *model.Job) {
			s.Publish(job)
		}, job))
	case cron:
		goJob, err = s.scheduler.NewJob(gocron.CronJob(*job.CronExpr, true), gocron.NewTask(func(job *model.Job) {
			s.Publish(job)
		}, job))
	}
	if err != nil {
		return err
	}
	if goJob == nil {
		return errors.New("create new job failed")
	}
	return nil
}

// Publish jobExecution event to queue
func (s *HegelScheduler) Publish(job *model.Job) error {

	if err := s.productor.Publish(job, Queue); err != nil {
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

// GetAvailableJobs get new jobs from db
func (s *HegelScheduler) GetAvailableJobs() ([]*model.Job, error) {
	s.jobRepo.GetAvailableJobs(context.Background(), nil)
	return nil, nil
}
