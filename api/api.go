package api

import "github.com/google/wire"

var ApiProviderSet = wire.NewSet(NewJobApi)
