package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	rerr := errors.New("Test error")
	cerr := New(ServerErr, "", rerr)
	assert.Equal(t, rerr.Error(), cerr.Error())
	terr := New(ClientErr, "New custom error", nil)
	assert.Equal(t, terr.Error(), "New custom error")
}
