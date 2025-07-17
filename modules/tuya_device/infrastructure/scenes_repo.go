package infrastructure

import (
	"encoding/json"
	"io"
	"os"

	"smart-home/modules/tuya_device/api"
	"smart-home/modules/tuya_device/constants"
	"smart-home/modules/tuya_device/externals"
)

type scenesRepo struct{}

func (s *scenesRepo) LoadScenes() (*api.Scenes, error) {
	file := constants.SCENES_FILE

	jsonFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	devices := &api.Scenes{}

	fileBytes, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, &devices)

	if err != nil {
		return nil, err
	}

	return devices, nil
}

var _ externals.ScenesRepo = (*scenesRepo)(nil)

func NewScenesRepo() *scenesRepo {
	return &scenesRepo{}
}
