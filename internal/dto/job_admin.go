package dto

import (
	"hegelscheduler/internal/model"
	"time"
)

type CreateJobRequest struct {
	Name          string        `gorm:"type:varchar(255);not null;comment:任务名称"`
	Description   string        `gorm:"type:text;comment:任务描述"`
	Type          string        `gorm:"type:varchar(64);not null;default:SIMPLE;comment:任务类型（SIMPLE/CRON/FIXED_DELAY）"`
	CronExpr      *string       `gorm:"type:varchar(128);comment:Cron表达式"`
	RunAt         *time.Time    `gorm:"comment:定时任务的运行时间"`
	RetryCount    int           `gorm:"default:0;comment:失败重试次数"`
	RetryInterval int           `gorm:"default:60;comment:重试间隔（秒）"`
	Timeout       int           `gorm:"default:300;comment:任务超时时间（秒）"`
	Payload       model.JSONMap `gorm:"type:json;comment:任务负载（worker 所需数据）"`
	TargetURL     string        `gorm:"type:varchar(512);not null;comment:任务执行的目标URL"`
	Method        string        `gorm:"type:varchar(8);default:POST;comment:请求方法"`
	Headers       model.JSONMap `gorm:"type:json;comment:请求头"`
}

type UpdateJobRequest struct {
	ID            uint64        `gorm:"primaryKey;autoIncrement;comment:任务ID"`
	Name          string        `gorm:"type:varchar(255);not null;comment:任务名称"`
	Description   string        `gorm:"type:text;comment:任务描述"`
	Type          string        `gorm:"type:varchar(64);not null;default:SIMPLE;comment:任务类型（SIMPLE/CRON/FIXED_DELAY）"`
	CronExpr      *string       `gorm:"type:varchar(128);comment:Cron表达式"`
	RunAt         *time.Time    `gorm:"comment:定时任务的运行时间"`
	RetryCount    int           `gorm:"default:0;comment:失败重试次数"`
	RetryInterval int           `gorm:"default:60;comment:重试间隔（秒）"`
	Timeout       int           `gorm:"default:300;comment:任务超时时间（秒）"`
	Payload       model.JSONMap `gorm:"type:json;comment:任务负载（worker 所需数据）"`
	TargetURL     string        `gorm:"type:varchar(512);not null;comment:任务执行的目标URL"`
	Method        string        `gorm:"type:varchar(8);default:POST;comment:请求方法"`
	Headers       model.JSONMap `gorm:"type:json;comment:请求头"`
}

type JobInfo struct {
	ID            uint64        `gorm:"primaryKey;autoIncrement;comment:任务ID"`
	Name          string        `gorm:"type:varchar(255);not null;comment:任务名称"`
	Description   string        `gorm:"type:text;comment:任务描述"`
	Type          string        `gorm:"type:varchar(64);not null;default:SIMPLE;comment:任务类型（SIMPLE/CRON/FIXED_DELAY）"`
	CronExpr      *string       `gorm:"type:varchar(128);comment:Cron表达式"`
	RunAt         *time.Time    `gorm:"comment:定时任务的运行时间"`
	RetryCount    int           `gorm:"default:0;comment:失败重试次数"`
	RetryInterval int           `gorm:"default:60;comment:重试间隔（秒）"`
	Timeout       int           `gorm:"default:300;comment:任务超时时间（秒）"`
	Payload       model.JSONMap `gorm:"type:json;comment:任务负载（worker 所需数据）"`
	TargetURL     string        `gorm:"type:varchar(512);not null;comment:任务执行的目标URL"`
	Method        string        `gorm:"type:varchar(8);default:POST;comment:请求方法"`
	Headers       model.JSONMap `gorm:"type:json;comment:请求头"`
	Status        string        `gorm:"type:varchar(32);not null;default:ENABLED;comment:任务状态"`
	CreatedBy     string        `gorm:"type:varchar(128);comment:创建人"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime"`
}
