package cmd

import (
	"bytes"
	"testing"
)

func TestAppCmd(t *testing.T) {
	var b bytes.Buffer
	cmd := appCmd
	cmd.SetOut(&b)
	cmd.SetArgs([]string{"--config", "../.log-aggregator_example.json"})
	cmd.Execute()

}
