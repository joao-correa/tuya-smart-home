package usecases

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"hash/crc32"
	"reflect"
	"strings"

	"smart-home/modules/tuya_device/api"
	"smart-home/modules/tuya_device/constants"
	"smart-home/modules/tuya_device/externals"
)

type translatedMessagePayload struct {
	DevId string         `json:"devId"`
	Uid   string         `json:"uid"`
	GwId  string         `json:"gwId"`
	Dps   map[string]any `json:"dps"`
}

type buildMessageUsecase struct {
	encryptMessage externals.EncryptMessage
}

func (b *buildMessageUsecase) BuildMessage(message *api.Message) ([]byte, error) {
	var payloadEncoded []byte
	var encryptedPayload []byte
	var err error

	if message.Payload != nil {
		payloadEncoded, err = json.Marshal(b.translateMessage(message.Payload))

		if err != nil {
			return nil, err
		}

		payloadEncoded = []byte(strings.ReplaceAll(string(payloadEncoded), " ", ""))

		encryptedPayload, err = b.encryptMessage.Encrypt(payloadEncoded, message.Key)
		if err != nil {
			return nil, err
		}
	}

	if !constants.NO_PROTOCOL_HEADER_CMDS[message.Cmd] {
		headerBytes := make([]byte, 12)
		cmdBytes := []byte(message.Version)
		cmdBytes = append(cmdBytes, headerBytes...)
		encryptedPayload = append(cmdBytes, encryptedPayload...)
	}

	messageLength := len(encryptedPayload) + 8

	bts := encryptedPayload

	// Write in seuqence
	// prefix, msgNumber, cmd, payload length, payload, checksum, suffix
	dataBuffer := new(bytes.Buffer)
	binary.Write(dataBuffer, binary.BigEndian, uint32(constants.PREFIX_55AA_VALUE))
	binary.Write(dataBuffer, binary.BigEndian, uint32(0))
	binary.Write(dataBuffer, binary.BigEndian, uint32(message.Cmd))
	binary.Write(dataBuffer, binary.BigEndian, uint32(messageLength))

	bts = append(dataBuffer.Bytes(), bts...)
	checksum := crc32.ChecksumIEEE(bts) & 0xFFFFFFFF

	dataBuffer.Reset()
	binary.Write(dataBuffer, binary.BigEndian, uint32(checksum))
	binary.Write(dataBuffer, binary.BigEndian, uint32(constants.SUFFIX_55AA_VALUE))

	return append(bts, dataBuffer.Bytes()...), nil
}

func (b *buildMessageUsecase) translateMessage(message *api.MessagePayload) translatedMessagePayload {

	dpsTranlation := map[string]string{
		"Switch":     "20",
		"Mode":       "21",
		"Brightness": "22",
		"ColourTemp": "23",
		"Colour":     "24",
		"Scene":      "25",
		"SceneData":  "25",
		"Timer":      "26",
		"Music":      "28",
		"ValueMin":   "10",
		"ValueMax":   "1000",
		"ValueHex":   "hsv16",
	}

	translated := translatedMessagePayload{
		DevId: message.DevId,
		Uid:   message.Uid,
		GwId:  message.GwId,
		Dps:   map[string]any{},
	}

	v := reflect.ValueOf(message.Dps)
	tp := v.Type()

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := tp.Field(i)

		// Ignore zero values if type is not boolean
		// On Booelan we need to keep zero values (false)
		if fieldType.Type.Kind() != reflect.Bool && field.Interface() == reflect.Zero(fieldType.Type).Interface() {
			continue
		}

		translated.Dps[dpsTranlation[fieldType.Name]] = field.Interface()
	}

	return translated
}

var _ api.BuildMessageUsecase = (*buildMessageUsecase)(nil)

func NewBuildMessageUsecase(
	encryptMessage externals.EncryptMessage,
) *buildMessageUsecase {
	return &buildMessageUsecase{
		encryptMessage: encryptMessage,
	}
}
