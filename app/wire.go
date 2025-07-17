//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	tuyaDevice "smart-home/modules/tuya_device"
	tuyaDeviceApi "smart-home/modules/tuya_device/api"
)

type Api struct {
	tuyaDeviceApi *tuyaDeviceApi.TuyaDeviceApi
}

func loadApis(
	tuyaDeviceApi *tuyaDeviceApi.TuyaDeviceApi,
) (*Api, error) {
	return &Api{
		tuyaDeviceApi,
	}, nil
}

func LoadApis() (*Api, func(), error) {
	wire.Build(
		loadApis,
		tuyaDevice.TuyaDeviceApiAllSet,
	)

	return &Api{}, nil, nil
}

