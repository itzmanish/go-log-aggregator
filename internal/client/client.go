package client

type Client interface {
	Init(opts ...Option) error
	Options() *Options
	Send(data interface{}) error
	Recv(out interface{}) error
	SendAndRecv(req interface{}, res interface{}) error
	Close() error
	String() string
}

func NewClient(opts ...Option) (Client, error) {
	return NewTcpClient(opts...)
}
