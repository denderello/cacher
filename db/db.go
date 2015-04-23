package db

type KeyValueDatabase interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
