package client

import (
	"context"
	"errors"
	"io"
	"net"
	"syscall"
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/codec"
	jc "github.com/itzmanish/go-loganalyzer/internal/codec/json"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
)

type tcpClient struct {
	opts Options
	conn net.Conn
	out  chan *codec.Packet
}

func (t *tcpClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	err := t.Connect()
	if err != nil {
		return err
	}
	go t.Read()
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

func (t *tcpClient) Options() Options {
	return t.opts
}

func (t *tcpClient) Out() chan *codec.Packet {
	return t.out
}

func (t *tcpClient) send(data interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	}
	return t.opts.Codec.Write(data)
}

func (t *tcpClient) Send(data interface{}) error {
	tsend := t.send
	ctx, cancel := context.WithTimeout(context.Background(), t.opts.Timeout)
	defer cancel()
	ch := make(chan error, t.opts.MaxRetries+1)
	var terr error

	for i := 0; i < t.opts.MaxRetries; i++ {
		go func() {
			ch <- tsend(data)
		}()

		select {
		case <-ctx.Done():
			return errors.New("timeout hit")

		case err := <-ch:
			// if the call succeeded lets bail early
			if err == nil {
				return nil
			}
			if errors.Is(err, syscall.EPIPE) {
				t.conn.Close()
				err = t.Connect()
			}
			terr = err
		}
	}

	return terr

}

// Recv is for recieving however it is not working
// Do not use this for now. Will fix it later.
func (t *tcpClient) Recv(out interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetReadDeadline(time.Now().Add(t.opts.Timeout))
	}
	return t.opts.Codec.Read(out)
	// return json.NewDecoder(t.conn).Decode(&out)
}

func (t *tcpClient) Read() error {
	for {
		var msg codec.Packet
		err := t.opts.Codec.Read(&msg)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				logger.Error(neterr)
				// return fmt.Errorf("TCP timeout : %s", err.Error())
			} else {
				logger.Errorf("Received error decoding message: %s", err.Error())
			}
		}

		t.out <- &msg
	}
}

func (t *tcpClient) Close() error {
	close(t.out)
	return t.opts.Codec.Close()
}

func (t *tcpClient) String() string {
	return "TCP Client"
}

func NewTcpClient(opts ...Option) (Client, error) {
	t := &tcpClient{
		out: make(chan *codec.Packet),
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
