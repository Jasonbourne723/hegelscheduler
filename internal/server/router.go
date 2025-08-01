package server

import (
	"github.com/gin-gonic/gin"
	"hegelscheduler/api"
)

type router struct {
	jobApi *api.JobApi
}

func NewRouter(jobApi *api.JobApi) *router {
	return &router{
		jobApi: jobApi,
	}
}

func (router *router) Router() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	return r
}
