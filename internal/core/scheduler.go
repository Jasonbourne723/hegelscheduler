package core

import (
	"errors"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"hegelscheduler/internal/model"
	"log"
	"time"
)

type HegelScheduler struct {
	IsLeader  bool
	jobChan   chan *model.Job
	stopChan  chan bool
	scheduler gocron.Scheduler
}

func NewHegelScheduler() *HegelScheduler {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	return &HegelScheduler{
		jobChan:   make(chan *model.Job, 100),
		stopChan:  make(chan bool),
		IsLeader:  false,
		scheduler: scheduler,
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
func (s *HegelScheduler) Start() error {
	go s.Poll()
	go s.Consumer()
	s.scheduler.Start()
	return nil
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
func (s *HegelScheduler) Publish(*model.Job) error {
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
				newJobs, err := s.GetNewJobs()
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

// GetNewJobs get new jobs from db
func (s *HegelScheduler) GetNewJobs() ([]*model.Job, error) {
	return nil, nil
}
