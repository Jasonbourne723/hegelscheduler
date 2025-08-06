package server

import (
	"github.com/gin-gonic/gin"
	"hegelscheduler/api"
)

type router struct {
	jobAdminApi    *api.JobAdminApi
	jobExectionApi *api.JobExectionApi
}

func NewRouter(jobAdminApi *api.JobAdminApi, jobExectionApi *api.JobExectionApi) *router {
	return &router{
		jobAdminApi:    jobAdminApi,
		jobExectionApi: jobExectionApi,
	}
}

func (router *router) Router() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	jobRouter := r.Group("JobAdmin")
	{
		jobRouter.POST("/", router.jobAdminApi.Create)
		jobRouter.PUT("/", router.jobAdminApi.Update)
		jobRouter.DELETE("/", router.jobAdminApi.Delete)
		jobRouter.GET("/Page", router.jobAdminApi.PageList)
	}

	jobExectionRouter := r.Group("JobExection")
	{
		jobExectionRouter.POST("Running/:id", router.jobExectionApi.SetRunning)
		jobExectionRouter.PUT("Success/:id", router.jobExectionApi.SetSuccess)
		jobExectionRouter.DELETE("Failed/:id", router.jobExectionApi.SetFailed)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	return r
}
