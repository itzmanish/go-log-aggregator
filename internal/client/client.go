package client

import "github.com/itzmanish/go-loganalyzer/internal/transport"

type Client interface {
	Init(opts ...Option) error
	Options() Options
	Out() chan *transport.Packet
	Send(data interface{}) error
	Recv(out interface{}) error
	Read() error
	Close() error
	String() string
}

func NewClient(opts ...Option) (Client, error) {
	t := &tcpClient{
		out: make(chan *transport.Packet),
	}
	err := t.Init(opts...)
	if err != nil {
		return nil, err
	}
	return t, nil
}
