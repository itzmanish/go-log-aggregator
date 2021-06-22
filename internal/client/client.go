package client

type Client interface {
	Init(opts ...Option) error
	Options() Options
	Send([]byte) error
	Recv(out chan []byte) error
	String() string
}

func NewClient(opts ...Option) Client {
	t := &tcpClient{}
	t.Init(opts...)
	return t
}
