package infrastructure

import (
	"smart-home/modules/tuya_device/externals"

	"github.com/google/wire"
)

var AllInfrastructure = wire.NewSet(
	NewDeviceConnection, wire.Bind(new(externals.DeviceConnection), new(*deviceConnection)),
	NewDevicesRepo, wire.Bind(new(externals.DevicesRepo), new(*devicesRepo)),
	NewScenesRepo, wire.Bind(new(externals.ScenesRepo), new(*scenesRepo)),
	NewEncryptMessage, wire.Bind(new(externals.EncryptMessage), new(*encryptMessage)),
	NewDecryptMessage, wire.Bind(new(externals.DecryptMessage), new(*decryptMessage)),
)
