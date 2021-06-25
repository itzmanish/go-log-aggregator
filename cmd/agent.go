package cmd

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/itzmanish/go-loganalyzer/internal/client"
	"github.com/itzmanish/go-loganalyzer/internal/codec"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/queue"
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
	cli, err := client.NewClient(
		client.WithAddress(serverConfig.Host+":"+serverConfig.Port),
		client.WithTimeout(5*time.Second),
	)
	if err != nil {
		logger.Fatal(err)
	}
	q := queue.NewQueue(cli)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	go SendLogs(w, cli, q)
	<-exit
	err = cli.Close()
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Shutting down agent...")
}

func SendLogs(w watcher.Watcher, client client.Client, q queue.Queue) {
	sent := sync.Map{}
	go func() {
		for v := range w.Result() {
			req := &codec.Packet{ID: uuid.New(), Cmd: "log", Body: &codec.LogBody{
				Name:      v.Name,
				Log:       v.Log,
				Tags:      v.Tags,
				Timestamp: v.Timestamp,
			}, Timestamp: time.Now()}
			err := client.Send(req)
			if err != nil {
				logger.Error(err)
				q.Push(req)
			} else {
				sent.Store(req.ID, req)
			}
		}
	}()
	for v := range client.Out() {
		_, ok := sent.Load(v.ID)
		if ok {
			if v.Ack {
				sent.Delete(v.ID)
			}
		}
		if v.Error != nil {
			logger.Error(v.Error)
		}
	}

}
