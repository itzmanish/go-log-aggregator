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
		cmd.SetArgs([]string{"--config", "../.loganalyzer.json"})
		cmd.Execute()
	})
	t.Run("TestServerCmd", func(t *testing.T) {
		cmd.SetArgs([]string{"--config", "../.loganalyzer.json", "server"})
		go func() {
			cmd.Execute()
		}()
		<-time.After(1 * time.Second)
	})

	t.Run("TestAgentCmd", func(t *testing.T) {
		cmd.SetArgs([]string{"--config", "../.loganalyzer.json", "agent"})
		go func() {
			cmd.Execute()
		}()
		f, err := os.OpenFile("../sample/log.txt", os.O_APPEND|os.O_WRONLY, 0755)
		assert.Nil(t, err)
		for i := 0; i < 2; i++ {
			n, err := f.WriteString("foo\n")
			assert.Nil(t, err)
			assert.NotZero(t, n)
		}
		f.Close()
		<-time.After(1 * time.Second)
		t.Log(b.ReadString('\n'))
	})
}
