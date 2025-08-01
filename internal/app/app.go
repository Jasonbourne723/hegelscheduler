package app

import "hegelscheduler/internal/server"

type App struct {
	httpServer *server.HttpServer
}

func NewApp(httpServer *server.HttpServer) *App {
	return &App{
		httpServer: httpServer,
	}
}

func (app *App) Start() {
	app.httpServer.Run()
}
