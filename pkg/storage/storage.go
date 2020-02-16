package storage

type Storage interface {
	Upload(file []byte) (string, error)
}
