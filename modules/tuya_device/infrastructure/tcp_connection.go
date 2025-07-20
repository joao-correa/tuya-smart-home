package infrastructure

import (
	"bytes"
	"encoding/binary"
	"net"
	"time"

	"smart-home/modules/tuya_device/constants"
	"smart-home/modules/tuya_device/externals"
)

type deviceConnection struct {
}

func (d *deviceConnection) SendMessageWithResponse(cmd *externals.MessageWithResponse) ([]byte, error) {
	conn, err := d.connect(cmd.IpAddress, cmd.Port)
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(cmd.Message)
	if err != nil {
		return nil, err
	}

	bts := make([]byte, 0)
	// Read until find prefix and suffix bytes which indicates the start and end of the message
	for {
		prefix := bytes.Index(bts, []byte(constants.PREFIX_BIN))
		suffix := bytes.Index(bts, []byte(constants.SUFFIX_BIN))

		if prefix != -1 && suffix != -1 {
			break
		}

		newBytes := make([]byte, 24)
		_, err = conn.Read(newBytes)
		if err != nil {
			return nil, err
		}

		bts = append(bts, newBytes...)
	}

	// Decode header to get payload size, after this we can extract the payload from
	// rest of message and return it, then usecase might be able to decode message
	// using key and right algorithm.
	//
	// header has 16 bytes in total
	// a sequence of 4 int32 values
	// that means: prefix, senquence-number, cmd, payload-size.
	headerLen := 16
	headerData := make([]rune, 4)
	payloadSizeIdx := 3
	_, err = binary.Decode(bts[:headerLen], binary.BigEndian, headerData)
	if err != nil {
		return nil, err
	}

	// The payload is composed of: payload, checksum, suffix
	// where last 8 bytes are checksum and suffix bytes and must
	// be removed from message, otherwise it decrypt will fail
	var payloadSize = int32(headerData[payloadSizeIdx])
	payloadData := make([]byte, payloadSize)
	_, err = binary.Decode(bts[headerLen:], binary.BigEndian, payloadData)
	if err != nil {
		return nil, err
	}

	// the payload structure is as follow:
	// ret-code (4 bytes), ...encrypted-payload, checksum (4 bytes), suffix (4 bytes)
	// We need to get rid of ret code, checksum and suffix bytes and keep only the
	// payload.
	suffixSize := 4
	checksumSize := 4
	retCodeSize := 4

	encryptedPayload := payloadData[retCodeSize : len(payloadData)-(checksumSize+suffixSize)]

	return encryptedPayload, nil
}

func (d *deviceConnection) SendMessageWithoutResponse(cmd *externals.MessageWithoutResponse) error {
	conn, err := d.connect(cmd.IpAddress, cmd.Port)
	defer conn.Close()

	if err != nil {
		return err
	}

	_, err = conn.Write(cmd.Message)
	if err != nil {
		return err
	}

	bytes := make([]byte, 4)
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
