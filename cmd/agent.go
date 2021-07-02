package cmd

import (
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-log-aggregator/config"
	"github.com/itzmanish/go-log-aggregator/internal/client"
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/codec/gob"
	"github.com/itzmanish/go-log-aggregator/internal/logger"
	"github.com/itzmanish/go-log-aggregator/internal/queue"
	"github.com/itzmanish/go-log-aggregator/internal/watcher"
	"github.com/itzmanish/go-log-aggregator/tool"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Log aggregator agent for collecting logs and sending to server.",
	Run: func(cmd *cobra.Command, args []string) {
		RunAgent(cmd, args)
	},
}

func init() {
	appCmd.AddCommand(agentCmd)
}

func RunAgent(cmd *cobra.Command, args []string) {
	agentConfig := config.AgentConfig{}
	err := config.Scan("agent", &agentConfig)
	if err != nil {
		logger.Error(err)
	}
	files := tool.FilterFileWatcher(agentConfig.Watchers)
	w := watcher.NewFileWatcher(files)
	w.Watch()
	defer w.Close()

	serverConfig := config.ServerConfig{}
	err = config.Scan("server", &serverConfig)
	if err != nil {
		logger.Fatal(err)
	}
	cli, err := client.NewClient(
		client.WithCodec(gob.NewGobCodec()),
		client.WithAddress(serverConfig.Host+":"+serverConfig.Port),
		client.WithTimeout(agentConfig.Timeout),
		client.WithMaxRetries(agentConfig.MaxRetries),
	)
	if err != nil {
		logger.Fatal(err)
	}
	q := queue.NewQueue(queue.WithClient(cli), queue.WithTimeInterval(agentConfig.QueueFlushInterval), queue.WithMaxQueueSize(agentConfig.MaxQueueSize))
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
	go func() {
		for v := range w.Result() {
			req := &codec.Packet{ID: uuid.New(), Cmd: "log", Body: &codec.LogBody{
				Name:      v.Name,
				Log:       v.Log,
				Tags:      v.Tags,
				Timestamp: v.Timestamp,
			}, Timestamp: time.Now()}
			go func(req *codec.Packet) {
				res := &codec.Packet{}
				err := client.SendAndRecv(req, res)
				if err != nil {
					logger.Error(err)
					q.Push(req)
				}
			}(req)
		}
	}()

}
