package server

type Server interface {
	Init(opts ...Option) error
	Options() Options
	Start() error
	Stop() error
	Closed() bool
	String() string
}

func NewServer(opts ...Option) Server {
	t := tcpServer{
		close: make(chan bool, 1),
	}
	t.Init(opts...)
	return &t
}
