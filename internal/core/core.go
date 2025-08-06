package core

import (
	"github.com/google/wire"
)

const (
	Queue = "hegel.job"
)

var CoreProviderSet wire.ProviderSet = wire.NewSet(
	NewJobAdminService,
	NewJobExectionService,
)
