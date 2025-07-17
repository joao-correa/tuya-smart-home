package externals

type MessageWithResponse struct {
	IpAddress string
	Port      string
	Key       string
	Message   []byte
}

type MessageWithoutResponse struct {
	IpAddress string
	Port      string
	Key       string
	Message   []byte
}

type DeviceConnection interface {
	SendMessageWithResponse(cmd *MessageWithResponse) (string, error)
	SendMessageWithoutResponse(cmd *MessageWithoutResponse) error
}
