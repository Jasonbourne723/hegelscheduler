package api

import (
	"github.com/gin-gonic/gin"
	"hegelscheduler/internal/core"
	"hegelscheduler/internal/dto"
	"net/http"
)

type JobAdminApi struct {
	srv *core.JobAdminService
}

func NewJobAdminApi(srv *core.JobAdminService) *JobAdminApi {
	return &JobAdminApi{
		srv: srv,
	}
}

// Create  a new job
func (s *JobAdminApi) Create(ctx *gin.Context) {
	var request dto.CreateJobRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := s.srv.Create(ctx.Request.Context(), request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Update  job information
func (s *JobAdminApi) Update(ctx *gin.Context) {
	var request dto.UpdateJobRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := s.srv.Update(ctx.Request.Context(), request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Delete Jobs by jobIds
func (s *JobAdminApi) Delete(ctx *gin.Context) {
	var ids []uint64
	if ctx.BindJSON(&ids) == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": ids})
	}
	if err := s.srv.Delete(ctx.Request.Context(), ids); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// PageList jobs
func (s *JobAdminApi) PageList(ctx *gin.Context) {
	var request dto.PageRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if response, err := s.srv.PageList(ctx.Request.Context(), request); err != nil {
		ctx.JSON(http.StatusInternalServerError, response)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
