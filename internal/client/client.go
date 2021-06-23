package client

type Client interface {
	Init(opts ...Option) error
	Options() Options
	Send(data interface{}) error
	Recv(out interface{}) error
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
