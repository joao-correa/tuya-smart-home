package infrastructure

import (
	"encoding/json"
	"io"
	"os"

	"smart-home/modules/tuya_device/api"
	"smart-home/modules/tuya_device/constants"
	"smart-home/modules/tuya_device/externals"
)

type devicesRepo struct{}

func (d *devicesRepo) LoadDevices() (*api.Devices, error) {
	file := constants.DEVICES_FILE

	jsonFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	fileBytes, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	devices := &api.Devices{}
	err = json.Unmarshal(fileBytes, &devices)

	if err != nil {
		return nil, err
	}

	return devices, nil

}

var _ externals.DevicesRepo = (*devicesRepo)(nil)

func NewDevicesRepo() *devicesRepo {
	return &devicesRepo{}
}
