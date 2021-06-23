package client

import (
	"io"
	"net"
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/transport"
)

type tcpClient struct {
	opts  Options
	conn  net.Conn
	codec Codec
}

func (t *tcpClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	conn, err := net.Dial("tcp", t.opts.Address)
	t.conn = conn
	t.codec = NewCodec(t.conn)
	return err
}

func (t *tcpClient) Options() Options {
	return t.opts
}

func (t *tcpClient) Send(data *transport.Packet) error {
	if t.opts.Timeout != 0 {
		t.conn.SetWriteDeadline(time.Now().Add(t.opts.Timeout))
	}
	return t.codec.Encode(data)
}

func (t *tcpClient) Recv(out chan transport.Packet) error {
	if t.opts.Timeout != 0 {
		t.conn.SetReadDeadline(time.Now().Add(t.opts.Timeout))
	}
	for {
		var msg transport.Packet
		err := t.codec.Decode(&msg)
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}

		out <- msg
	}
}

func (t *tcpClient) String() string {
	return "TCP Client"
}
