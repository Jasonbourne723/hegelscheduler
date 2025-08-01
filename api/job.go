package api

import "hegelscheduler/internal/core"

type JobApi struct {
	srv *core.JobService
}

func NewJobApi(srv *core.JobService) *JobApi {
	return &JobApi{
		srv: srv,
	}
}
