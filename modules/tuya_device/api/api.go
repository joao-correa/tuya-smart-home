package api

type TuyaDeviceApi struct {
	BuildMessageUsecase
	ApplySceneUsecase
}

type BuildMessageUsecase interface {
	BuildMessage(message *Message) ([]byte, error)
}

type ApplySceneUsecase interface {
	ApplyScene(sceneName string) error
}

type Message struct {
	Cmd      int            `json:"cmd"`
	DeviceId string         `json:"deviceId"`
	Version  string         `json:"version"`
	Key      string         `json:"key"`
	Payload  MessagePayload `json:"payload"`
}

type MessagePayload struct {
	DevId string `json:"devId"`
	Uid   string `json:"uid"`
	GwId  string `json:"gwId"`
	Dps   Dps    `json:"dps"`
}

func NewTuyaDeviceApi(
	applySceneUsecase ApplySceneUsecase,
	buildMessageUsecase BuildMessageUsecase,
) *TuyaDeviceApi {
	return &TuyaDeviceApi{
		ApplySceneUsecase:   applySceneUsecase,
		BuildMessageUsecase: buildMessageUsecase,
	}
}
