package cmd

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func runTestServer(t *testing.T, addr chan string) error {
	l, err := net.Listen("tcp", ":33555")
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

func TestAgentCmd(t *testing.T) {
	addr := make(chan string, 1)
	runTestServer(t, addr)
	var b bytes.Buffer
	cmd := appCmd
	cmd.SetOut(&b)
	cmd.SetArgs([]string{"--config", "../.log-aggregator_example.json", "agent", "server.port", "33555"})
	go func() {
		err := cmd.Execute()
		assert.Nil(t, err)
	}()
	<-time.After(100 * time.Millisecond)
	f, err := os.OpenFile("../sample/log.txt", os.O_APPEND|os.O_WRONLY, 0755)
	assert.Nil(t, err)
	for i := 0; i < 20; i++ {
		n, err := f.WriteString("foo\n")
		assert.Nil(t, err)
		assert.NotZero(t, n)
	}
	f.Close()
	t.Log(b.ReadString('\n'))
}
