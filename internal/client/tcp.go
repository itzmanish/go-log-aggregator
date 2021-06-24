package client

import (
	"io"
	"net"
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/codec"
	jc "github.com/itzmanish/go-loganalyzer/internal/codec/json"
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

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return err
	}
	go t.Read()
	return err
}

func (t *tcpClient) Options() Options {
	return t.opts
}

func (t *tcpClient) Out() chan *codec.Packet {
	return t.out
}

func (t *tcpClient) Send(data interface{}) error {
	if t.opts.Timeout != 0 {
		t.conn.SetWriteDeadline(time.Now().Add(t.opts.Timeout))
	}
	return t.opts.Codec.Write(data)
	// db, err := json.Marshal(data)
	// if err != nil {
	// 	return err
	// }
	// _, err = t.conn.Write(db)
	// return err
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
			return err
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
			Codec: jc.NewCodec(),
		},
	}
	err := t.Init(opts...)
	if err != nil {
		return nil, err
	}
	return t, nil
}
