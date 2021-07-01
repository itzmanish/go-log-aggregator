package handler

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	s, err := store.NewFileStore(store.WithDirectory("../sample"))
	assert.Nil(t, err)
	h := NewHandler(s, 1*time.Second, 2)
	id := uuid.New()
	res, err := h.Handle(&codec.Packet{ID: id, Cmd: "log"})
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, id, res.ID)
	assert.Equal(t, true, res.Ack)
}
