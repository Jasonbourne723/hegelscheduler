package main

// @title           Job Scheduler API
// @version         1.0
// @description     分布式任务调度系统 API 文档
// @host            localhost:8080
// @BasePath        /
func main() {

	app := InitializeEvent()

	app.Start()

}
