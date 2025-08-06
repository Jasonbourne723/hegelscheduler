package api

import (
	"github.com/gin-gonic/gin"
	"hegelscheduler/internal/core"
	"net/http"
	"strconv"
)

type JobExectionApi struct {
	srv core.JobExectionService
}

func NewJobExectionApi(srv core.JobExectionService) *JobExectionApi {
	return &JobExectionApi{
		srv: srv,
	}
}

func (api *JobExectionApi) SetRunning(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	err := api.srv.SetRunning(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"status": http.StatusOK,
		})
	}
}

func (api *JobExectionApi) SetFailed(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	err := api.srv.SetFailed(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"status": http.StatusOK,
		})
	}
}

func (api *JobExectionApi) SetSuccess(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	err := api.srv.SetSuccess(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"status": http.StatusOK,
		})
	}
}
