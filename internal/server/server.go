package server

type Server interface {
	Init(opts ...Option) error
	Options() Options
	Start() error
	Stop() error
	Closed() bool
	String() string
}
