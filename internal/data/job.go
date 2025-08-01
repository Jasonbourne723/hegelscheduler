package data

import (
	"context"
	"gorm.io/gorm"
	"hegelscheduler/internal/model"
)

type JobRepo struct {
	*gorm.DB
}

func NewJobRepo(db *gorm.DB) *JobRepo {
	return &JobRepo{
		db,
	}
}

func (s *JobRepo) Create(ctx context.Context, job model.Job) error {

	if err := s.WithContext(ctx).Create(&job).Error; err != nil {
		return err
	}
	return nil
}

// Update  job information
func (s *JobRepo) Update(ctx context.Context, job model.Job) error {

	if err := s.WithContext(ctx).Save(&job).Error; err != nil {
		return err
	}
	return nil
}

// Delete Jobs by jobIds
func (s *JobRepo) Delete(ctx context.Context, ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}
	if err := s.WithContext(ctx).Delete(&model.Job{}, "id in (?)", ids).Error; err != nil {
		return err
	}
	return nil
}

// PageList jobs
func (s *JobRepo) PageList(ctx context.Context, index int, size int) (int64, []model.Job, error) {
	var (
		total int64
		list  []model.Job
	)
	query := s.WithContext(ctx).Model(&model.Job{})
	if err := query.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	if err := query.Find(&list).Offset(size * (index - 1)).Limit(size).Error; err != nil {
		return 0, nil, err
	}
	return total, list, nil
}
