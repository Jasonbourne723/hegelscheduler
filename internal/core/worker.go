package core

import "hegelscheduler/internal/config"

type Worker struct{}

func NewWorker(bs *config.BootStrap) *Worker {
	return &Worker{}
}
