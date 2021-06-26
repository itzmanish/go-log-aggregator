package tool

import (
	"os"

	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
)

// FilterFileWatcher filters for file watcher and returns only watcher which are type of filesystem
func FilterFileWatcher(watchers config.Watchers) config.Watchers {
	filewatchers := config.Watchers{}
	for _, watcher := range watchers {
		for _, tag := range watcher.Tags {
			if tag["key"] == "Type" && tag["value"] == "filesystem" {
				filewatchers = append(filewatchers, watcher)
			}
		}
	}
	return filewatchers
}

// GetSeekInfo gets the current file size
func GetSeekInfo(name string) int64 {
	info, err := os.Stat(name)
	if err != nil {
		logger.Error(err)
		return 0
	}
	return info.Size()
}
