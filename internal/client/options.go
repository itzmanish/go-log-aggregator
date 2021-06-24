package client

import (
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/codec"
)

type Options struct {
	Address    string
	MaxRetries int
	Timeout    time.Duration
	Codec      codec.Codec
}

type Option func(o *Options)

func WithAddress(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

func WithMaxRetries(retries int) Option {
	return func(o *Options) {
		o.MaxRetries = retries
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func WithCodec(c codec.Codec) Option {
	return func(o *Options) {
		o.Codec = c
	}
}
