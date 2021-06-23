package client

import (
	"encoding/json"
	"io"
	"net"
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/transport"
)

type tcpClient struct {
	opts Options
	conn net.Conn
	out  chan *transport.Packet
}

func (t *tcpClient) Init(opts ...Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	conn, err := net.Dial("tcp", t.opts.Address)
	t.conn = conn
	go t.Read()
	return err
}

func (t *tcpClient) Options() Options {
	return t.opts
}

func (t *tcpClient) Out() chan *transport.Packet {
	return t.out
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

// Recv is for recieving however it is not working
// Do not use this for now. Will fix it later.
func (t *tcpClient) Recv(out interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetReadDeadline(time.Now().Add(t.opts.Timeout))
	}
	return json.NewDecoder(t.conn).Decode(&out)
}

func (t *tcpClient) Read() error {
	decoder := json.NewDecoder(t.conn)
	for {
		var msg transport.Packet
		err := decoder.Decode(&msg)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		t.out <- &msg
	}
}

func (t *tcpClient) Close() error {
	close(t.out)
	return t.conn.Close()
}

func (t *tcpClient) String() string {
	return "TCP Client"
}
