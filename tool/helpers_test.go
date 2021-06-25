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

func TestGetSeekInfo(t *testing.T) {
	t.Run("Success with correct path", func(t *testing.T) {
		n := GetSeekInfo("../sample/log.txt")
		assert.Equal(t, n, int64(31))
	})
	t.Run("Fail with wrong path", func(t *testing.T) {
		n := GetSeekInfo("sample/log.txt")
		assert.Equal(t, n, int64(0))
	})
}
