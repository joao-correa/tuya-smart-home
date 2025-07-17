package externals

import (
	"smart-home/modules/tuya_device/api"
)

type DevicesRepo interface {
	LoadDevices() (*api.Devices, error)
}
