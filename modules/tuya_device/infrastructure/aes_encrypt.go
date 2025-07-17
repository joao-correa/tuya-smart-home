package infrastructure

import (
	"crypto/aes"

	"smart-home/modules/tuya_device/externals"
)

type encryptMessage struct{}

func (e *encryptMessage) Encrypt(message []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	message = e.fixPadding(message, aes.BlockSize)

	encrypted := make([]byte, len(message))

	for i := 0; i < len(message); i += aes.BlockSize {
		block.Encrypt(encrypted[i:i+aes.BlockSize], message[i:i+aes.BlockSize])
	}

	return encrypted, nil
}

func (e *encryptMessage) fixPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

var _ externals.EncryptMessage = (*encryptMessage)(nil)

func NewEncryptMessage() *encryptMessage {
	return &encryptMessage{}
}
