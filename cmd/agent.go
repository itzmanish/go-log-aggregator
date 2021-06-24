package cmd

import (
	"os"
	"os/signal"
	"time"

	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/itzmanish/go-loganalyzer/internal/client"
	"github.com/itzmanish/go-loganalyzer/internal/codec"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/watcher"
	"github.com/itzmanish/go-loganalyzer/tool"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Log analyzer agent for collecting logs and sending to server.",
	Run: func(cmd *cobra.Command, args []string) {
		RunAgent(cmd, args)
	},
}

func init() {
	appCmd.AddCommand(agentCmd)
}

func RunAgent(cmd *cobra.Command, args []string) {
	watchers := config.Watchers{}
	err := config.Scan("watchers", &watchers)
	if err != nil {
		logger.Error(err)
	}
	files := tool.FilterFileWatcher(watchers)
	w := watcher.NewFileWatcher(files)
	w.Watch()
	defer w.Close()

	serverConfig := config.ServerConfig{}
	err = config.Scan("server", &serverConfig)
	if err != nil {
		logger.Fatal(err)
	}
	cli, err := client.NewClient(client.WithAddress(serverConfig.Host+":"+serverConfig.Port), client.WithTimeout(5*time.Second))
	if err != nil {
		logger.Fatal(err)
	}
	go func() {
		for v := range w.Result() {
			err = cli.Send(&codec.Packet{ID: "1", Cmd: "log", Body: &codec.LogBody{
				Name:      v.Name,
				Log:       v.Log,
				Tags:      v.Tags,
				Timestamp: v.Timestamp,
			}, Timestamp: time.Now()})
			if err != nil {
				logger.Error(err)
			}
		}
	}()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	go func() {
		for v := range cli.Out() {
			logger.Info(*v)
		}
	}()

	<-exit
	err = cli.Close()
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Shutting down agent...")
}
