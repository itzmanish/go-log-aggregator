package watcher

import "os"

type Watcher interface {
	Watch()
}

type watch struct {
	file os.File
}

func (w *watch) Watcher() {

}
