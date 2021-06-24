package store

import (
	"github.com/philippgille/gokv"
)

type Store interface {
	gokv.Store
	String() string
}

var Stores map[string]func(opts ...Option) (Store, error) = map[string]func(opts ...Option) (Store, error){
	"file": NewFileStore,
	"s3":   NewS3Store,
}

var defaultStore, _ = NewFileStore()

func Get(key string, value interface{}) (bool, error) {
	return defaultStore.Get(key, value)
}

func Set(key string, value interface{}) error {
	return defaultStore.Set(key, value)
}

func Delete(key string) error {
	return defaultStore.Delete(key)
}

func Close() error {
	return defaultStore.Close()
}

func String() string {
	return defaultStore.String()
}
