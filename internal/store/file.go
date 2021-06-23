package store

import (
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/file"
)

type fileStore struct {
	opts  Options
	store gokv.Store
}

func (f *fileStore) Set(key string, value interface{}) error {
	return f.store.Set(key, value)
}

func (f *fileStore) Get(key string, value interface{}) (bool, error) {
	return f.store.Get(key, value)
}

func (f *fileStore) Delete(key string) error {
	return f.store.Delete(key)
}

func (f *fileStore) Close() error {
	return f.store.Close()
}

func (f *fileStore) String() string {
	return "File Store"
}

func NewFileStore(opts ...Option) (Store, error) {
	fs := &fileStore{}
	for _, o := range opts {
		o(&fs.opts)
	}
	foptions := file.DefaultOptions
	if len(fs.opts.Directory) != 0 {
		foptions.Directory = fs.opts.Directory
	}
	s, err := file.NewStore(foptions)
	if err != nil {
		return nil, err
	}
	fs.store = s
	return fs, nil
}
