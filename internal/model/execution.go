package model

import "time"

type Execution struct {
	Id            int64
	JobId         int64
	Name          string
	JobType       int
	Status        int
	Host          string
	StartTime     time.Time
	EndTime       time.Time
	SchedulerTime time.Time
	CreatedAt     time.Time
}
