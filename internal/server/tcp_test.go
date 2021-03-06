package server

import (
	"encoding/json"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/stretchr/testify/assert"
)

type testHandler struct{}

func (th *testHandler) Handle(req *codec.Packet) (*codec.Packet, error) {
	return req, nil
}

func TestTCPServer(t *testing.T) {
	port := "34253"
	server := NewServer()
	err := server.Init(WithPort(port), WithHandler(&testHandler{}))
	assert.Nil(t, err)
	assert.Equal(t, server.Options().Port, port)
	go func() {
		time.Sleep(1 * time.Second)
		conn, err := net.Dial("tcp4", ":"+port)
		assert.Nil(t, err)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		go func() {
			out := codec.Packet{}
			err := json.NewDecoder(conn).Decode(&out)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, "log", out.Cmd)
			conn.Close()
		}()
		err = json.NewEncoder(conn).Encode(&codec.Packet{ID: uuid.New(), Cmd: "log"})
		assert.Nil(t, err)
		server.Stop()
	}()
	assert.Equal(t, "TCP server", server.String())
	err = server.Start()
	assert.Nil(t, err)
}
