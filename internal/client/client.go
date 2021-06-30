package client

import "github.com/itzmanish/go-log-aggregator/internal/codec"

type Client interface {
	Init(opts ...Option) error
	Options() Options
	Out() chan *codec.Packet
	Send(data interface{}) error
	Recv(out interface{}) error
	Read()
	Close() error
	String() string
}

func NewClient(opts ...Option) (Client, error) {
	return NewTcpClient(opts...)
}
