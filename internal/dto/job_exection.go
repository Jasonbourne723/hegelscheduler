package dto

import (
	"hegelscheduler/internal/model"
)

type JobExectionDto struct {
	JobExectionId uint64        `json:"jobExectionId"`
	JobId         uint64        `json:"jobId"`
	Name          string        `gorm:"type:varchar(255);not null;comment:任务名称"`
	Description   string        `gorm:"type:text;comment:任务描述"`
	RetryCount    int           `gorm:"default:0;comment:失败重试次数"`
	RetryInterval int           `gorm:"default:60;comment:重试间隔（秒）"`
	Timeout       int           `gorm:"default:300;comment:任务超时时间（秒）"`
	Payload       model.JSONMap `gorm:"type:json;comment:任务负载（worker 所需数据）"`
	TargetURL     string        `gorm:"type:varchar(512);not null;comment:任务执行的目标URL"`
	Method        string        `gorm:"type:varchar(8);default:POST;comment:请求方法"`
	Headers       model.JSONMap `gorm:"type:json;comment:请求头"`
}
