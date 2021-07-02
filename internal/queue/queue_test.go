package queue

import (
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-log-aggregator/internal/client"
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/stretchr/testify/assert"
)

func runTestServer(t *testing.T, addr chan string) error {
	l, err := net.Listen("tcp", ":5541")
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
			db, err := ioutil.ReadAll(conn)
			assert.Nil(t, err)
			conn.Write(db)
			return
		}
	}()
	return nil
}

func TestQueue(t *testing.T) {
	addr := make(chan string, 1)
	err := runTestServer(t, addr)
	if err != nil {
		t.Fatal(err)
	}
	c, err := client.NewClient(client.WithAddress(<-addr))
	assert.Nil(t, err)
	q := NewQueue(WithClient(c), WithTimeInterval(1*time.Millisecond), WithMaxQueueSize(5))
	assert.Equal(t, "Memory queue", q.String())
	id := uuid.New()
	t.Run("TestPush", func(t *testing.T) {
		q.Push(&codec.Packet{
			ID:  id,
			Cmd: "log",
			Ack: true,
		})
	})
	<-time.After(5 * time.Millisecond)
	t.Run("TestGet", func(t *testing.T) {
		v, ok := q.Get(q.Length())
		assert.Equal(t, true, ok)
		assert.Equal(t, &codec.Packet{
			ID:  id,
			Cmd: "log",
			Ack: true,
		}, v)
	})
	t.Run("TestPop", func(t *testing.T) {
		q.Pop(q.Length())
	})
	t.Run("TestOptions", func(t *testing.T) {
		assert.Equal(t, q.Options().Client, c)

	})
	t.Run("FailWithMaxCapacity0", func(t *testing.T) {
		q.Init(WithMaxQueueSize(0))
		q.Push(&codec.Packet{
			ID:  id,
			Cmd: "log",
			Ack: true,
		})
	})
	assert.Equal(t, 0, q.Length())
}
