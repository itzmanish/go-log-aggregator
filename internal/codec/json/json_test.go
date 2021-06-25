package json

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
	jc := NewCodec()
	assert.Equal(t, "json", jc.String())

	t.Run("TestInit", func(t *testing.T) {
		var b CloseableBuffer
		jc.Init(&b)
	})

	t.Run("TestWrite", func(t *testing.T) {
		err := jc.Write(map[string]string{"foo": "bar"})
		assert.Nil(t, err)
	})

	t.Run("TestRead", func(t *testing.T) {
		var out map[string]string
		err := jc.Read(&out)
		assert.Nil(t, err)
		assert.Equal(t, map[string]string{"foo": "bar"}, out)
	})
	t.Run("TestWriteNil", func(t *testing.T) {
		err := jc.Write(nil)
		assert.Nil(t, err)
	})
	t.Run("TestReadNil", func(t *testing.T) {
		err := jc.Read(nil)
		assert.Nil(t, err)
	})
	t.Run("TestClose", func(t *testing.T) {
		err := jc.Close()
		assert.Nil(t, err)
	})
}
