package gob

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CloseableBuffer struct {
	bytes.Buffer
}

func (b *CloseableBuffer) Close() error {
	return nil
}

func TestJson(t *testing.T) {
	gc := NewGobCodec()
	assert.Equal(t, "gob", gc.String())

	t.Run("TestInit", func(t *testing.T) {
		var b CloseableBuffer
		gc.Init(&b)
	})

	t.Run("TestWrite", func(t *testing.T) {
		err := gc.Write(map[string]string{"foo": "bar"})
		assert.Nil(t, err)
	})

	t.Run("TestRead", func(t *testing.T) {
		var out map[string]string
		err := gc.Read(&out)
		assert.Nil(t, err)
		assert.Equal(t, map[string]string{"foo": "bar"}, out)
	})
	t.Run("TestWriteNil", func(t *testing.T) {
		err := gc.Write(nil)
		assert.NotNil(t, err)
	})
	t.Run("TestReadNil", func(t *testing.T) {
		err := gc.Read(nil)
		assert.NotNil(t, err)
	})
	t.Run("TestClose", func(t *testing.T) {
		err := gc.Close()
		assert.Nil(t, err)
	})
}
