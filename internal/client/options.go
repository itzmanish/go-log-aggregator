package client

import "time"

type Options struct {
	Address    string
	MaxRetries int32
	Timeout    time.Duration
}

type Option func(o *Options)

func WithAddress(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

func WithMaxRetries(retries int32) Option {
	return func(o *Options) {
		o.MaxRetries = retries
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}
