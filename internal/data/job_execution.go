package data

import "gorm.io/gorm"

type JobExecutionRepo struct {
	db *gorm.DB
}

func NewJobExecutionRepo(db *gorm.DB) *JobExecutionRepo {
	return &JobExecutionRepo{
		db: db,
	}
}
