package api

import "github.com/google/wire"

var TuayDeviceApi = wire.NewSet(
	NewTuyaDeviceApi,
)
