package client

import (
	"encoding/json"
	"net"
	"time"
)

type tcpClient struct {
	opts Options
	conn net.Conn
}

func (t *tcpClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	conn, err := net.Dial("tcp", t.opts.Address)
	t.conn = conn
	return err
}

func (t *tcpClient) Options() Options {
	return t.opts
}

func (t *tcpClient) Send(data interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetWriteDeadline(time.Now().Add(t.opts.Timeout))
	}
	db, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = t.conn.Write(db)
	return err
}

func (t *tcpClient) Recv(out interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetReadDeadline(time.Now().Add(t.opts.Timeout))
	}
	return json.NewDecoder(t.conn).Decode(&out)
}

func (t *tcpClient) String() string {
	return "TCP Client"
}
