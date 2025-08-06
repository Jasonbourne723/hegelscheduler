package data

import (
	"gorm.io/gorm"
	"hegelscheduler/internal/model"
)

type JobExecutionRepo struct {
	db *gorm.DB
}

func NewJobExecutionRepo(db *gorm.DB) *JobExecutionRepo {
	return &JobExecutionRepo{
		db: db,
	}
}

func (r *JobExecutionRepo) Create(jobExecution *model.JobExecution) error {
	return r.db.Create(jobExecution).Error
}

func (r *JobExecutionRepo) SetRunning(id uint64) error {
	return r.db.Model(&model.JobExecution{}).
		Where("id = ? and status = ?", id, model.JobExecutionStatusReady).
		Updates(map[string]interface{}{"status": model.JobExecutionStatusRunning}).Error
}

func (r *JobExecutionRepo) SetSuccess(id uint64) error {
	return r.db.Model(&model.JobExecution{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": model.JobExecutionStatusRunning}).Error
}

func (r *JobExecutionRepo) SetFailed(id uint64) error {
	return r.db.Model(&model.JobExecution{}).
		Where("id = ? and status != ?", id, model.JobExecutionStatusSuccess).
		Updates(map[string]interface{}{"status": model.JobExecutionStatusFailed}).Error
}
