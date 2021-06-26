package store

import (
	"github.com/philippgille/gokv"
)

// Store interface provide methods that are required in a store
type Store interface {
	gokv.Store
	String() string
}

// Stores is the list of available store
var Stores map[string]func(opts ...Option) (Store, error) = map[string]func(opts ...Option) (Store, error){
	"file": NewFileStore,
	"s3":   NewS3Store,
}

var defaultStore, _ = NewFileStore()

// Get a value from store using key and store in value parameter.
// It return boolean value of value found or not and a error
func Get(key string, value interface{}) (bool, error) {
	return defaultStore.Get(key, value)
}

// Set a key value in store and returns a error
func Set(key string, value interface{}) error {
	return defaultStore.Set(key, value)
}

// Delete value from store using key
func Delete(key string) error {
	return defaultStore.Delete(key)
}

// Close the store
func Close() error {
	return defaultStore.Close()
}

// String return the current store name
func String() string {
	return defaultStore.String()
}
