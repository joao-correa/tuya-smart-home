package externals

type EncryptMessage interface {
	Encrypt(message []byte, key string) ([]byte, error)
}
