package store

type Options struct {
	Directory string
}

type Option func(o *Options)

func WithDirectory(dir string) Option {
	return func(o *Options) {
		o.Directory = dir
	}
}
