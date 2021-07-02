package cmd

import (
	"bytes"
	"testing"
	"time"
)

func TestServerCmd(t *testing.T) {
	var b bytes.Buffer
	cmd := appCmd
	cmd.SetOut(&b)
	cmd.SetArgs([]string{"--config", "../.log-aggregator_example.json", "server"})
	go func() {
		cmd.Execute()
	}()
	<-time.After(10 * time.Millisecond)
}
