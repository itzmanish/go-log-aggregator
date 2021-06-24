package server

import "github.com/itzmanish/go-loganalyzer/internal/codec"

type Options struct {
	Port    string
	Host    string
	Handler Handler
	Codec   codec.Codec
}

type Option func(o *Options)

func WithHost(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

func WithPort(port string) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func WithHandler(handler Handler) Option {
	return func(o *Options) {
		o.Handler = handler
	}
}

func WithCodec(c codec.Codec) Option {
	return func(o *Options) {
		o.Codec = c
	}
}
