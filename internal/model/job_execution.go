package model

import (
	"time"
)

type JobExecution struct {
	ID            uint64     `gorm:"primaryKey;autoIncrement;comment:执行记录ID"`
	JobID         uint64     `gorm:"not null;index:idx_job_id;comment:关联的任务ID"`
	ScheduledTime time.Time  `gorm:"not null;index:idx_scheduled_time;comment:任务计划运行时间"`
	StartTime     *time.Time `gorm:"comment:实际开始时间"`
	EndTime       *time.Time `gorm:"comment:实际结束时间"`
	Status        string     `gorm:"type:varchar(32);not null;index:idx_status;comment:状态（RUNNING/SUCCESS/FAILED/...）"`
	Result        string     `gorm:"type:text;comment:执行结果或错误信息"`
	WorkerID      string     `gorm:"type:varchar(128);comment:worker实例ID"`
	WorkerIP      string     `gorm:"type:varchar(128);comment:workerIP地址"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
}

const (
	JobExecutionStatusReady   = "ready"
	JobExecutionStatusRunning = "running"
	JobExecutionStatusSuccess = "success"
	JobExecutionStatusFailed  = "failed"
)
