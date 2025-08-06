package core

import (
	"hegelscheduler/internal/config"
	"hegelscheduler/internal/data"
)

type JobExectionService struct {
	repo *data.JobExecutionRepo
	bs   *config.BootStrap
}

func NewJobExectionService(bs *config.BootStrap, repo *data.JobExecutionRepo) *JobExectionService {
	return &JobExectionService{
		repo: repo,
		bs:   bs,
	}
}

func (s *JobExectionService) SetRunning(id uint64) error {
	return s.repo.SetRunning(id)
}

func (s *JobExectionService) SetFailed(id uint64) error {
	return s.repo.SetFailed(id)
}

func (s *JobExectionService) SetSuccess(id uint64) error {
	return s.repo.SetSuccess(id)
}
