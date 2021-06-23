package watcher

import (
	"time"

	"github.com/hpcloud/tail"
	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/tool"
)

type Result struct {
	Name string `json:"name"`
	// TODO: right now tags should be sent every time will fix it later for performance improvements
	Tags      []map[string]interface{} `json:"tags"`
	Log       string                   `json:"log"`
	Timestamp time.Time                `json:"time"`
}

type fileWatcher struct {
	files  config.Watchers
	result chan Result
}

func NewFileWatcher(files config.Watchers) Watcher {
	return &fileWatcher{
		files:  files,
		result: make(chan Result),
	}
}

func (fw *fileWatcher) Watch() {
	for _, file := range fw.files {
		go func(file config.Watcher) {
			t, err := tail.TailFile(file.Watch, tail.Config{Follow: true, Location: &tail.SeekInfo{
				Offset: tool.GetSeekInfo(file.Watch),
			}})
			if err != nil {
				logger.Error(err)
			}
			for line := range t.Lines {
				res := Result{Name: file.Watch, Tags: file.Tags, Log: line.Text, Timestamp: line.Time}
				fw.result <- res
			}
		}(file)
	}
}

func (fw *fileWatcher) Result() chan Result {
	return fw.result
}

func (fw *fileWatcher) Close() {
	close(fw.result)
}

func (fw *fileWatcher) String() string {
	return "File watcher"
}
