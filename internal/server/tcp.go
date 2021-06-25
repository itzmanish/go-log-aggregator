package server

import (
	"errors"
	"io"
	"net"
	"sync/atomic"

	"github.com/google/uuid"
	"github.com/itzmanish/go-loganalyzer/internal/codec"
	"github.com/itzmanish/go-loganalyzer/internal/codec/json"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
)

type tcpServer struct {
	opts     Options
	listener net.Listener
	close    chan bool
	closed   int32
}

// Init initialize tcp server
func (t *tcpServer) Init(opts ...Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	return nil
}

// Options return options of tcp server
func (t *tcpServer) Options() Options {
	return t.opts
}

func (t *tcpServer) Closed() bool {
	return atomic.LoadInt32(&t.closed) > 0
}

func (t *tcpServer) Start() error {
	logger.Infof("Server is running on %s:%s", t.opts.Host, t.opts.Port)
	l, err := net.Listen("tcp4", t.opts.Host+":"+t.opts.Port)
	if err != nil {
		return err
	}
	t.listener = l
	defer t.listener.Close()
	for {
		c, err := t.listener.Accept()
		if err != nil {
			select {
			case <-t.close:
				return nil
			default:
				logger.Error("error on connection with client: ", err)
			}
		}
		go t.handleConnection(c)

	}
}

func (t *tcpServer) Stop() error {
	logger.Info("Server is stopping...")
	if t.Closed() {
		return nil
	}
	if !atomic.CompareAndSwapInt32(&t.closed, 0, 1) {
		return errors.New("unable to stop server")
	}
	close(t.close)
	return t.listener.Close()
}

func (*tcpServer) String() string {
	return "TCP server"
}

func (t *tcpServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	t.opts.Codec.Init(conn)
	for {
		if t.Closed() {
			return
		}
		var msg codec.Packet
		err := t.opts.Codec.Read(&msg)
		if err != nil {
			if err != io.EOF {
				logger.Error("read error", err)
			}
			return
		}
		if msg.ID != uuid.Nil && t.opts.Handler != nil {
			res, err := t.opts.Handler.Handle(&msg)
			if err != nil {
				err = t.opts.Codec.Write(&codec.Packet{
					Error: err,
				})
				if err != nil {
					logger.Error(err)
				}
				return

			}
			err = t.opts.Codec.Write(res)
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}
}

func NewTcpServer(opts ...Option) Server {
	t := tcpServer{
		close: make(chan bool, 1),
		opts: Options{
			Codec: json.NewCodec(),
		},
	}
	t.Init(opts...)
	return &t
}
