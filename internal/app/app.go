package app

import (
	"gorm.io/gorm"
	"hegelscheduler/internal/model"
	"hegelscheduler/internal/server"
)

type App struct {
	httpServer *server.HttpServer
	db         *gorm.DB
}

func NewApp(httpServer *server.HttpServer, db *gorm.DB) *App {
	return &App{
		httpServer: httpServer,
		db:         db,
	}
}

func (app *App) Start() {
	app.autoMigrate()
	app.httpServer.Run()
}

func (app *App) autoMigrate() {
	app.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.Job{},
		&model.JobExecution{},
	)
}
