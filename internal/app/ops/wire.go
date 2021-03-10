package ops

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewApp, NewOptions)
