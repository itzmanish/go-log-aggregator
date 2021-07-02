package cmd

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAppCmd(t *testing.T) {
	var b bytes.Buffer
	cmd := appCmd
	cmd.SetOut(&b)
	t.Run("TestWithoutArgs", func(t *testing.T) {
		cmd.SetArgs([]string{"--config", "../.log-aggregator_example.json"})
		cmd.Execute()
	})
	t.Run("TestServerCmd", func(t *testing.T) {
		cmd.SetArgs([]string{"--config", "../.log-aggregator_example.json", "server"})
		go func() {
			cmd.Execute()
		}()
		<-time.After(100 * time.Millisecond)
	})

	t.Run("TestAgentCmd", func(t *testing.T) {
		cmd.SetArgs([]string{"--config", "../.log-aggregator_example.json", "agent"})
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
	})
}
