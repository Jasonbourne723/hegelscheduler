package core

import "github.com/go-co-op/gocron/v2"

type Scheduler struct {
	IsLeader bool
	Id       int64
}

func (s Scheduler) Election() {

}

func (s Scheduler) Start() error {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	scheduler.NewJob(gocron.CronJob("", false), gocron.NewTask(func() {}))

	scheduler.Start()
	return nil
}
