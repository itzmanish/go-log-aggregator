package watcher

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"
	"github.com/itzmanish/go-logent/pkg/config"
	"github.com/itzmanish/go-logent/pkg/logger"
)

type Watcher interface {
	Watch()
}

type watch struct {
	file os.File
}

func Watch() {
	watchers := []config.Watcher{}
	err := config.Scan("watchers", &watchers)
	if err != nil {
		logger.Log(err)
	}
	t, err := tail.TailFile(watchers[0].Watch, tail.Config{Follow: true})
	logger.Log(err)
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
