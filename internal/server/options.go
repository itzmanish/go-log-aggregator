package server

import "io"

type Handler func(c io.ReadWriteCloser)

type Options struct {
	Port    string
	Host    string
	Handler Handler
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
