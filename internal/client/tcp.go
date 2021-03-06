package client

import (
	"context"
	"errors"
	"net"
	"syscall"
	"time"

	jc "github.com/itzmanish/go-log-aggregator/internal/codec/json"
)

type tcpClient struct {
	opts Options
	conn net.Conn
}

func (t *tcpClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	err := t.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (t *tcpClient) Connect() error {
	conn, err := net.Dial("tcp", t.opts.Address)
	if err != nil {
		return err
	}
	t.conn = conn
	t.opts.Codec.Init(conn)
	err = conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return err
	}

	return conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
}

func (t *tcpClient) Options() *Options {
	return &t.opts
}

func (t *tcpClient) send(data interface{}) error {
	return t.opts.Codec.Write(data)
}

func (t *tcpClient) Send(data interface{}) error {
	ch := make(chan error, t.opts.MaxRetries)
	var terr error
	send := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), t.opts.Timeout)
		defer cancel()

		go func(ch chan error) {
			ch <- t.send(data)
		}(ch)

		select {
		case <-ctx.Done():
			return errors.New("timeout hit")

		case err := <-ch:
			return err
		}
	}
	for i := 0; i <= t.opts.MaxRetries; i++ {
		err := send()
		// if the call succeeded lets bail early
		if err == nil {
			return nil
		}
		if e, ok := err.(*net.OpError); ok && e.Err.Error() == "use of closed network connection" || errors.Is(err, syscall.EPIPE) {
			t.conn.Close()
			err = t.Connect()
		}

		terr = err
	}
	return terr
}

// Recv is for receiving however it is not working
// Do not use this for now. Will fix it later.
func (t *tcpClient) Recv(out interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetReadDeadline(time.Now().Add(t.opts.Timeout))
	}
	return t.opts.Codec.Read(out)
}

func (t *tcpClient) SendAndRecv(req interface{}, res interface{}) error {
	err := t.Send(req)
	if err != nil {
		return err
	}
	err = t.Recv(res)
	return err
}

func (t *tcpClient) Close() error {
	return t.opts.Codec.Close()
}

func (t *tcpClient) String() string {
	return "TCP Client"
}

func NewTcpClient(opts ...Option) (Client, error) {
	t := &tcpClient{
		opts: Options{
			Codec:      jc.NewCodec(),
			MaxRetries: 2,
		},
	}
	err := t.Init(opts...)
	if err != nil {
		return nil, err
	}
	return t, nil
}
