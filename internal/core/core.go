package core

import (
	"github.com/google/wire"
)

var CoreProviderSet wire.ProviderSet = wire.NewSet(
	NewJobService,
)
