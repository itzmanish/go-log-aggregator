package tool

import "github.com/itzmanish/go-logent/config"

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
