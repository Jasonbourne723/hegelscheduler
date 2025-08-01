package data

import "gorm.io/gorm"

type JobRepo struct {
	db *gorm.DB
}

func NewJobRepo(db *gorm.DB) *JobRepo {
	return &JobRepo{
		db: db,
	}
}
