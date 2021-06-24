package client

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-loganalyzer/internal/codec"
	"github.com/stretchr/testify/assert"
)

func runMockServer(t *testing.T) net.Listener {
	l, err := net.Listen("tcp", ":32111")
	assert.Nil(t, err)
	go func() {
		for {
			c, err := l.Accept()
			assert.Nil(t, err)
			defer c.Close()
			for {
				select {
				case <-time.After(5 * time.Second):
					return
				default:
					for {
						time.Sleep(100 * time.Millisecond)
						n, err := c.Write([]byte("Hi"))
						assert.Nil(t, err)
						assert.Greater(t, n, 0)
					}
				}
			}
		}
	}()
	return l
}

func TestTCPClient(t *testing.T) {
	l := runMockServer(t)
	defer l.Close()
	var client Client

	t.Run("TestNewClient", func(t *testing.T) {
		client, err := NewClient(WithAddress(l.Addr().String()))
		assert.NotNil(t, client)
		assert.Nil(t, err)
	})

	t.Run("TestInit", func(t *testing.T) {
		err := client.Init(WithMaxRetries(2), WithTimeout(5*time.Second))
		assert.Nil(t, err)
	})

	t.Run("TestOptions", func(t *testing.T) {
		opt := client.Options()
		assert.Equal(t, int32(2), opt.MaxRetries)
	})

	t.Run("TestString", func(t *testing.T) {
		assert.Equal(t, "TCP Client", client.String())
	})

	t.Run("TestSend", func(t *testing.T) {
		data := &codec.Packet{
			ID: uuid.New(),
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
