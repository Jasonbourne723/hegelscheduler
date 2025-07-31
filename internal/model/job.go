package model

import "time"

type Job struct {
	Id        int64
	Name      string
	JobType   int
	Status    int
	Cron      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
