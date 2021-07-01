package server

import (
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/handler"
)

type Options struct {
	Port    string
	Host    string
	Handler []handler.Handler
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

func WithHandler(handler handler.Handler) Option {
	return func(o *Options) {
		o.Handler = append(o.Handler, handler)
	}
}

func WithCodec(c codec.Codec) Option {
	return func(o *Options) {
		o.Codec = c
	}
}
