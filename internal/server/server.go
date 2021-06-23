package server

import (
	"io"

	"github.com/itzmanish/go-loganalyzer/internal/transport"
)

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
	Handle(req *transport.Packet, w io.Writer) error
}

func NewServer(opts ...Option) Server {
	t := tcpServer{
		close: make(chan bool, 1),
	}
	t.Init(opts...)
	return &t
}
