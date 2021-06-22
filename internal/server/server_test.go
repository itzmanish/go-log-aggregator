package server

import (
	"bufio"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTCPServer(t *testing.T) {
	port := "34253"
	server := NewServer()
	err := server.Init(WithPort(port))
	assert.Nil(t, err)
	assert.Equal(t, server.Options().Port, port)
	go func() {
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp4", ":"+port)
		assert.Nil(t, err)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		defer conn.Close()
		go func() {
			out, err := bufio.NewReader(conn).ReadString('\n')
			assert.Nil(t, err)
			assert.Equal(t, "Hi", out)
		}()
		n, err := conn.Write([]byte("Hi"))
		assert.Nil(t, err)
		assert.Greater(t, n, 0)
		err = server.Stop()
		assert.Nil(t, err)
	}()
	err = server.Start()
	assert.Nil(t, err)
	assert.Equal(t, "TCP server", server.String())
}
