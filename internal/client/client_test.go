package client

import (
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-loganalyzer/internal/codec"
	"github.com/stretchr/testify/assert"
)

func runTestServer(t *testing.T, addr chan string) error {
	l, err := net.Listen("tcp", ":5544")
	if !assert.Nil(t, err) {
		return err
	}
	assert.NotNil(t, l)
	addr <- l.Addr().String()
	go func() {
		conn, err := l.Accept()
		assert.Nil(t, err)
		defer conn.Close()
		for {
			b, err := ioutil.ReadAll(conn)
			assert.Nil(t, err)
			n, err := conn.Write(b)
			assert.Nil(t, err)
			assert.NotZero(t, n)
			return
		}
	}()
	return nil
}

func TestTCPClient(t *testing.T) {
	var client Client
	addr := make(chan string, 1)
	err := runTestServer(t, addr)
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.New()
	t.Run("TestNewClient", func(t *testing.T) {
		client, err = NewClient(WithAddress(<-addr))
		assert.Nil(t, err)
	})

	t.Run("TestInit", func(t *testing.T) {
		err := client.Init(WithMaxRetries(2), WithTimeout(5*time.Second))
		assert.Nil(t, err)
	})

	t.Run("TestOptions", func(t *testing.T) {
		opt := client.Options()
		assert.Equal(t, 2, opt.MaxRetries)
	})

	t.Run("TestString", func(t *testing.T) {
		assert.Equal(t, "TCP Client", client.String())
	})

	t.Run("TestSend", func(t *testing.T) {
		data := &codec.Packet{
			ID: id,
		}
		err := client.Send(data)
		assert.Nil(t, err)
	})

	t.Run("TestRecv", func(t *testing.T) {
		out := make(chan codec.Packet)
		go func() {
			client.Recv(out)
			close(out)
		}()
		for {
			select {
			case <-time.After(5 * time.Second):
				return
			case res, ok := <-out:
				if ok {
					log.Println(res)
				}
			}
		}
	})

}
