package core

import (
	"context"
	"hegelscheduler/internal/config"
	"hegelscheduler/internal/data"
	"hegelscheduler/internal/dto"
	"hegelscheduler/internal/mapper"
	"hegelscheduler/internal/model"
)

type JobAdminService struct {
	repo *data.JobRepo
	bs   *config.BootStrap
}

func NewJobAdminService(repo *data.JobRepo, bs *config.BootStrap) *JobAdminService {
	return &JobAdminService{
		repo: repo,
		bs:   bs,
	}
}

// Create  a new job
func (s *JobAdminService) Create(ctx context.Context, req dto.CreateJobRequest) error {

	var job model.Job
	if err := mapper.Map(&job, &req); err != nil {
		return err
	}
	return s.repo.Create(ctx, job)
}

// Update  job information
func (s *JobAdminService) Update(ctx context.Context, req dto.UpdateJobRequest) error {

	var job model.Job
	if err := mapper.Map(&job, &req); err != nil {
		return err
	}
	return s.repo.Update(ctx, job)
}

// Delete Jobs by jobIds
func (s *JobAdminService) Delete(ctx context.Context, ids []uint64) error {

	return s.repo.Delete(ctx, ids)
}

// PageList jobs
func (s *JobAdminService) PageList(ctx context.Context, req dto.PageRequest) (*dto.PageResponse[dto.JobInfo], error) {
	total, list, err := s.repo.PageList(ctx, req.Index, req.Size)
	if err != nil {
		return nil, err
	}
	items := make([]*dto.JobInfo, len(list))
	if err := mapper.Map(&list, &items); err != nil {
		return nil, err
	}
	return &dto.PageResponse[dto.JobInfo]{
		total,
		items,
	}, nil
}
