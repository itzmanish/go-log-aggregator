package client

import (
	"bufio"
	"fmt"
	"io"
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

func (t *tcpClient) Send(data []byte) error {
	if t.opts.Timeout != 0 {
		t.conn.SetWriteDeadline(time.Now().Add(t.opts.Timeout))
	}
	_, err := t.conn.Write(data)
	return err
}

func (t *tcpClient) Recv(out chan []byte) error {
	if t.opts.Timeout != 0 {
		t.conn.SetReadDeadline(time.Now().Add(t.opts.Timeout))
	}
	for {
		buffer, err := bufio.NewReader(t.conn).ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return err
			}
		}
		if len(buffer) < 1 {
			continue
		}
		fmt.Println(buffer)
		out <- buffer
	}
}

func (t *tcpClient) String() string {
	return "TCP Client"
}
