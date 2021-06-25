package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileStore(t *testing.T) {
	fs, err := NewFileStore(WithDirectory("../../sample"))
	assert.Nil(t, err)
	assert.NotNil(t, fs)
	assert.Equal(t, "File Store", fs.String())

	t.Run("Test Set to store", func(t *testing.T) {
		err := fs.Set("foo", "bar")
		assert.Nil(t, err)
	})
	t.Run("Test Get from store", func(t *testing.T) {
		var out string
		found, err := fs.Get("foo", &out)
		assert.Nil(t, err)
		assert.Equal(t, true, found)
		assert.Equal(t, "bar", out)
	})
	t.Run("Test Delete from store", func(t *testing.T) {
		err := fs.Delete("foo")
		assert.Nil(t, err)
	})
	t.Run("Test Close file store", func(t *testing.T) {
		err := fs.Close()
		assert.Nil(t, err)
	})
}
