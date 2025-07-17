package externals

type DecryptMessage interface {
	Decrypt(message []byte, key string) ([]byte, error)
}
