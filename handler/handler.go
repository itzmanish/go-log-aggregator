package handler

import (
	"time"

	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/handler"
	"github.com/itzmanish/go-log-aggregator/internal/logger"
	"github.com/itzmanish/go-log-aggregator/internal/store"
)

type srvHandler struct {
	store         store.Store
	buffer        chan *codec.Packet
	flushInterval time.Duration
}

func NewHandler(s store.Store, interval time.Duration, maxChunkCap int) handler.Handler {
	h := &srvHandler{
		store:         s,
		buffer:        make(chan *codec.Packet, maxChunkCap),
		flushInterval: interval,
	}
	go h.Flush()
	return h
}

func (h *srvHandler) Flush() {
	ticker := time.NewTicker(h.flushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logger.Info("Auto flushing after ", h.flushInterval)
			h.flush()
		default:
			if len(h.buffer) == cap(h.buffer) {
				logger.Info("Buffer full. Flushing now.")
				h.flush()
			}
		}
	}
}

func (h *srvHandler) flush() {
	data := []*codec.Packet{}
	length := len(h.buffer)
	for i := 0; i < length; i++ {
		data = append(data, <-h.buffer)
	}
	if len(data) == 0 {
		return
	}
	err := h.store.Set(time.Now().String(), data)
	if err != nil {
		logger.Error(err)
	}
}

func (h *srvHandler) Handle(req *codec.Packet) (*codec.Packet, error) {
	switch req.Cmd {
	case "log":
		logger.Info("Got your request: ", req.Body)
		h.buffer <- req
		ack := &codec.Packet{
			ID:        req.ID,
			Ack:       true,
			Timestamp: time.Now(),
		}
		return ack, nil

	default:
		logger.Info(req)
	}
	return nil, nil
}
