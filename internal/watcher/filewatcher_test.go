package watcher

import (
	"os"
	"testing"
	"time"

	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/stretchr/testify/assert"
)

func TestWatcher(t *testing.T) {
	watcherConfig := config.Watchers{
		config.Watcher{Watch: "../../sample/log.txt",
			Tags: []map[string]interface{}{
				{
					"Name": "Test Watcher",
				}},
		},
	}
	var w Watcher
	t.Run("Create new watcher", func(t *testing.T) {
		w = NewFileWatcher(watcherConfig)
		assert.NotNil(t, w)
		assert.Equal(t, "File watcher", w.String())
	})

	t.Run("Watch for changes", func(t *testing.T) {
		go func() {
			w.Watch()
			<-time.After(time.Second * 1)
			w.Close()
		}()
		f, err := os.OpenFile(watcherConfig[0].Watch, os.O_APPEND|os.O_WRONLY, 0755)
		assert.Nil(t, err)
		for i := 0; i < 2; i++ {
			n, err := f.WriteString("foo\n")
			assert.Nil(t, err)
			assert.NotZero(t, n)
		}
		f.Close()
		for v := range w.Result() {
			assert.Equal(t, "foo", v.Log)
			assert.Equal(t, watcherConfig[0].Watch, v.Name)
		}
	})

}
