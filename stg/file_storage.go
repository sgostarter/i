package stg

type FileStorage interface {
	WriteFile(name string, d []byte) error
	ReadFile(name string) ([]byte, error)
}
