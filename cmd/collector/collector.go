package collector

import (
	"github.com/itzmanish/go-logent/config"
	"github.com/itzmanish/go-logent/internal/logger"
	"github.com/itzmanish/go-logent/internal/watcher"
	"github.com/itzmanish/go-logent/tool"
	"github.com/spf13/cobra"
)

func RunCollector(cmd *cobra.Command, args []string) {
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
