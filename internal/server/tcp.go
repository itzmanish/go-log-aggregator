package server

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"sync/atomic"

	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/transport"
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
	decoder := json.NewDecoder(conn)
	for {
		select {
		case <-t.close:
			return
		default:
			var msg transport.Packet
			err := decoder.Decode(&msg)
			if err != nil {
				if err != io.EOF {
					logger.Error("read error", err)
				}
				return
			}
			if msg.ID != "" && t.opts.Handler != nil {
				err = t.opts.Handler.Handle(&msg, conn)
				if err != nil {
					logger.Error(err)
					return
				}
			}
		}
	}
}
