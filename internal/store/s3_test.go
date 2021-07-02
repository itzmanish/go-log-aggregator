package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3Store(t *testing.T) {
	s3, err := NewS3Store(WithDirectory("s3test"), WithPathStyleAddressing(true), WithS3Endpoint(os.Getenv("S3_ENDPOINT")),
		WithAWSAccessKey(os.Getenv("AWS_ACCESS_KEY")), WithAWSSecretAccessKey(os.Getenv("AWS_SECRET_KEY")),
	)
	assert.Nil(t, err)
	if !assert.NotNil(t, s3) {
		t.Fatal(err)
	}
	assert.Equal(t, "S3 Store", s3.String())

	t.Run("Test Set to store", func(t *testing.T) {
		err := s3.Set("foo", "bar")
		assert.Nil(t, err)
	})
	t.Run("Test Get from store", func(t *testing.T) {
		var out string
		found, err := s3.Get("foo", &out)
		assert.Nil(t, err)
		assert.Equal(t, true, found)
		assert.Equal(t, "bar", out)
	})
	t.Run("Test Delete from store", func(t *testing.T) {
		err := s3.Delete("foo")
		assert.Nil(t, err)
	})
	t.Run("Test Close file store", func(t *testing.T) {
		err := s3.Close()
		assert.Nil(t, err)
	})
}
