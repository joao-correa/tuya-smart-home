package infrastructure

import (
	"fmt"
	"net"
	"time"

	"smart-home/modules/tuya_device/externals"
)

type deviceConnection struct{}

func (d *deviceConnection) SendMessageWithResponse(cmd *externals.MessageWithResponse) (string, error) {
	conn, err := d.connect(cmd.IpAddress, cmd.Port)
	defer conn.Close()

	if err != nil {
		return "", err
	}

	r, err := conn.Write(cmd.Message)

	if err != nil {
		return "", err
	}

	fmt.Printf("%v", r)

	bytes := make([]byte, 4024)
	_, err = conn.Read(bytes)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (d *deviceConnection) SendMessageWithoutResponse(cmd *externals.MessageWithoutResponse) error {
	conn, err := d.connect(cmd.IpAddress, cmd.Port)
	defer conn.Close()

	if err != nil {
		return err
	}

	r, err := conn.Write(cmd.Message)

	if err != nil {
		return err
	}

	fmt.Printf("%v", r)

	bytes := make([]byte, 4024)
	_, err = conn.Read(bytes)

	return err
}

func (d *deviceConnection) connect(ipAddress string, Port string) (net.Conn, error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(ipAddress, Port))

	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(5 * time.Second))

	return conn, nil
}

var _ externals.DeviceConnection = (*deviceConnection)(nil)

func NewDeviceConnection() *deviceConnection {
	return &deviceConnection{}
}
