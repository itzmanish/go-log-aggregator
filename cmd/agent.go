package cmd

import (
	"github.com/itzmanish/go-loganalyzer/config"
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
	for v := range w.Result() {
		logger.Info(v)
	}

}
