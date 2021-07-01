package server

// Server
type Server interface {
	Init(opts ...Option) error
	Options() Options
	Start() error
	Stop() error
	Closed() bool
	String() string
}

func NewServer(opts ...Option) Server {
	return NewTcpServer(opts...)
}
