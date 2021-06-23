package client

import "github.com/itzmanish/go-loganalyzer/internal/transport"

type Client interface {
	Init(opts ...Option) error
	Options() Options
	Send(data *transport.Packet) error
	Recv(out chan transport.Packet) error
	String() string
}

func NewClient(opts ...Option) (Client, error) {
	t := &tcpClient{}
	err := t.Init(opts...)
	if err != nil {
		return nil, err
	}
	return t, nil
}
