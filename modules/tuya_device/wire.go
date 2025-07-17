package tuya_device

import (
	"github.com/google/wire"

	"smart-home/modules/tuya_device/api"
	"smart-home/modules/tuya_device/infrastructure"
	"smart-home/modules/tuya_device/usecases"
)

var TuyaDeviceApiAllSet = wire.NewSet(
	usecases.AllUsecases,
	infrastructure.AllInfrastructure,
	api.TuayDeviceApi,
)
