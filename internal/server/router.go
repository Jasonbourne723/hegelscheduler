package server

import (
	"github.com/gin-gonic/gin"
	"hegelscheduler/api"
)

type router struct {
	jobAdminApi *api.JobAdminApi
}

func NewRouter(jobAdminApi *api.JobAdminApi) *router {
	return &router{
		jobAdminApi: jobAdminApi,
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	return r
}
