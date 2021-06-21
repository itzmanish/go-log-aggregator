package tool

import (
	"testing"

	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/stretchr/testify/assert"
)

func TestFilterFileWatcher(t *testing.T) {
	watchers := config.Watchers{
		config.Watcher{
			Watch: "test/sfome.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "filesystem",
				},
			},
		},
		config.Watcher{
			Watch: "test/other.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "service",
				},
			},
		},
		config.Watcher{
			Watch: "test/fs2.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "filesystem",
				},
			},
		},
	}
	expectedWatcher := config.Watchers{
		config.Watcher{
			Watch: "test/sfome.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "filesystem",
				},
			},
		},
		config.Watcher{
			Watch: "test/fs2.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "filesystem",
				},
			},
		},
	}
	actualWatchers := FilterFileWatcher(watchers)
	assert.Equal(t, expectedWatcher, actualWatchers)
}
