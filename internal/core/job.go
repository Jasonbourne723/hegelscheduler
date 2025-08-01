package core

import "hegelscheduler/internal/data"

type JobService struct {
	repo *data.JobRepo
}

func NewJobService(repo *data.JobRepo) *JobService {
	return &JobService{
		repo: repo,
	}
}

func (s *JobService) Create() error {
	return nil
}

func (s *JobService) Update() error {
	return nil
}

func (s *JobService) Delete() error {
	return nil
}

func (s *JobService) PageList() error {
	return nil
}
