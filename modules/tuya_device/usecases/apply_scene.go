package usecases

import (
	"fmt"

	"smart-home/modules/tuya_device/api"
	"smart-home/modules/tuya_device/constants"
	"smart-home/modules/tuya_device/externals"
)

type applySceneUsecase struct {
	scenesRepo          externals.ScenesRepo
	devicesRepo         externals.DevicesRepo
	deviceConnection    externals.DeviceConnection
	buildMessageUsecase api.BuildMessageUsecase
}

func (a *applySceneUsecase) ApplyScene(sceneName string) error {
	scenes, err := a.scenesRepo.LoadScenes()

	if err != nil {
		return err
	}

	scene, ok := scenes.Scenes[sceneName]
	if !ok {
		return fmt.Errorf("scene not found")
	}

	devicesOnScene := map[string]bool{}
	for _, deviceID := range scene.DeviceIds {
		devicesOnScene[deviceID] = true
	}

	devices, err := a.devicesRepo.LoadDevices()

	if err != nil {
		return err
	}

	for _, device := range devices.Devices {
		if _, ok := devicesOnScene[device.ID]; !ok {
			continue
		}

		message := api.Message{
			Cmd:      constants.COMMAND_TYPE_CONTROL,
			Version:  "3.3",
			Key:      device.Key,
			DeviceId: device.ID,
			Payload: api.MessagePayload{
				DevId: device.ID,
				Uid:   device.ID,
				GwId:  device.ID,
				Dps:   scene.Dps,
			},
		}

		messageBytes, err := a.buildMessageUsecase.BuildMessage(&message)

		if err != nil {
			return err
		}

		err = a.deviceConnection.SendMessageWithoutResponse(&externals.MessageWithoutResponse{
			IpAddress: device.IP,
			Port:      "6668",
			Key:       device.Key,
			Message:   messageBytes,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

var _ api.ApplySceneUsecase = (*applySceneUsecase)(nil)

func NewApplySceneUsecase(
	scenesRepo externals.ScenesRepo,
	devicesRepo externals.DevicesRepo,
	deviceConnection externals.DeviceConnection,
	buildMessageUsecase api.BuildMessageUsecase,
) *applySceneUsecase {
	return &applySceneUsecase{
		scenesRepo:          scenesRepo,
		devicesRepo:         devicesRepo,
		deviceConnection:    deviceConnection,
		buildMessageUsecase: buildMessageUsecase,
	}
}
