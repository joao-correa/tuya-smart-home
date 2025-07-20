package infrastructure

import (
	"crypto/aes"

	"smart-home/modules/tuya_device/externals"
)

type decryptMessage struct{}

func (e *decryptMessage) Decrypt(message []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(message))

	// Decrypt each 16-byte block
	for i := 0; i < len(message); i += aes.BlockSize {
		block.Decrypt(decrypted[i:i+aes.BlockSize], message[i:i+aes.BlockSize])
	}

	return decrypted, nil
}

var _ externals.DecryptMessage = (*decryptMessage)(nil)

func NewDecryptMessage() *decryptMessage {
	return &decryptMessage{}
}
