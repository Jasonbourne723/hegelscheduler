//go:build wireinject

package main

import (
	"github.com/google/wire"
	"hegelscheduler/api"
	"hegelscheduler/internal/app"
	"hegelscheduler/internal/core"
	"hegelscheduler/internal/data"
	"hegelscheduler/internal/server"
)

func InitializeEvent() *app.App {
	wire.Build(api.ApiProviderSet, data.ProviderSet, core.CoreProviderSet, server.ServerProviderSet, app.NewApp)
	return &app.App{}
}
