package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	assert.Equal(t, "File Store", String())

	t.Run("Test Set to store", func(t *testing.T) {
		err := Set("foo", "bar")
		assert.Nil(t, err)
	})
	t.Run("Test Get from store", func(t *testing.T) {
		var out string
		found, err := Get("foo", &out)
		assert.Nil(t, err)
		assert.Equal(t, true, found)
		assert.Equal(t, "bar", out)
	})
	t.Run("Test Delete from store", func(t *testing.T) {
		err := Delete("foo")
		assert.Nil(t, err)
	})
	t.Run("Test Close file store", func(t *testing.T) {
		err := Close()
		assert.Nil(t, err)
	})
}
