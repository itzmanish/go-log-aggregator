package tool

import (
	"testing"

	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/stretchr/testify/assert"
)

func TestFilterFileWatcher(t *testing.T) {
	watchers := config.Watchers{
		config.Watcher{
			Watch: "sample/log.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "filesystem",
				},
			},
		},
		config.Watcher{
			Watch: "sample/other.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "service",
				},
			},
		},
		config.Watcher{
			Watch: "sample/log2.txt",
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
			Watch: "sample/log.txt",
			Tags: []map[string]interface{}{
				{
					"key":   "Type",
					"value": "filesystem",
				},
			},
		},
		config.Watcher{
			Watch: "sample/log2.txt",
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
		assert.NotZero(t, n)
	})
	t.Run("Fail with wrong path", func(t *testing.T) {
		n := GetSeekInfo("sample/log.txt")
		assert.Zero(t, n)
	})
}
