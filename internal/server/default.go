package server

import (
	"bufio"
	"errors"
	"io"
	"net"
	"sync/atomic"

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
	for {
		select {
		case <-t.close:
			return
		default:
			buffer, err := bufio.NewReader(conn).ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					logger.Error("read error", err)
					return
				}
			}
			if len(buffer) < 1 {
				return
			}
			logger.Info("Client message:", string(buffer[:len(buffer)-1]))
			conn.Write(buffer)
		}
	}
}

func NewServer(opts ...Option) Server {
	t := tcpServer{
		close: make(chan bool, 1),
	}
	t.Init(opts...)
	return &t
}
