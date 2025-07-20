package usecases

import (
	"smart-home/modules/tuya_device/api"
	"smart-home/modules/tuya_device/constants"
	"smart-home/modules/tuya_device/externals"
)

type getDevicesStatus struct {
	devicesRepo         externals.DevicesRepo
	deviceConnection    externals.DeviceConnection
	buildMessageUsecase api.BuildMessageUsecase
	decryptMessage      externals.DecryptMessage
}

func (a *getDevicesStatus) GetDevicesStatus() ([]byte, error) {
	devices, err := a.devicesRepo.LoadDevices()

	if err != nil {
		return nil, err
	}

	for _, device := range devices.Devices {
		message := api.Message{
			Cmd:      constants.COMMAND_TYPE_DP_QUERY,
			Version:  "3.3",
			Key:      device.Key,
			DeviceId: device.ID,
			Payload: &api.MessagePayload{
				DevId: device.ID,
				Uid:   device.ID,
				GwId:  device.ID,
			},
		}

		messageBytes, err := a.buildMessageUsecase.BuildMessage(&message)

		if err != nil {
			return nil, err
		}

		responseBytes, err := a.deviceConnection.SendMessageWithResponse(&externals.MessageWithResponse{
			IpAddress: device.IP,
			Port:      "6668",
			Key:       device.Key,
			Message:   messageBytes,
		})

		if err != nil {
			return nil, err
		}

		print("\n len response ", len(responseBytes))
		decodedBytes, err := a.decryptMessage.Decrypt(responseBytes, device.Key)
		if err != nil {
			return nil, err
		}

		print("\n len decoded ", len(decodedBytes))

		// TODO: translate from dps numbers to dps names
		// TODO: append to devices list
		// TODO: return list

		print("\n decode message ")
		print(string(decodedBytes))
		print("\n\n")
	}

	return nil, nil
}

var _ api.GetDevicesStatusUsecase = (*getDevicesStatus)(nil)

func NewGetDevicesStatus(
	devicesRepo externals.DevicesRepo,
	deviceConnection externals.DeviceConnection,
	buildMessageUsecase api.BuildMessageUsecase,
	decryptMessage externals.DecryptMessage,
) *getDevicesStatus {
	return &getDevicesStatus{
		devicesRepo:         devicesRepo,
		deviceConnection:    deviceConnection,
		buildMessageUsecase: buildMessageUsecase,
		decryptMessage:      decryptMessage,
	}
}
