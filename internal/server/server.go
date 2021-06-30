package server

import "github.com/itzmanish/go-log-aggregator/internal/codec"

// Server
type Server interface {
	Init(opts ...Option) error
	Options() Options
	Start() error
	Stop() error
	Closed() bool
	String() string
}

// Handler handles request from client
type Handler interface {
	Handle(in *codec.Packet) (*codec.Packet, error)
}

func NewServer(opts ...Option) Server {
	return NewTcpServer(opts...)
}
