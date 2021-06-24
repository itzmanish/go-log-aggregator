package handler

import (
	"time"

	"github.com/itzmanish/go-loganalyzer/internal/codec"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/server"
	"github.com/itzmanish/go-loganalyzer/internal/store"
)

type srvHandler struct {
	store store.Store
}

func NewHandler(s store.Store) server.Handler {
	return &srvHandler{
		store: s,
	}
}

func (h *srvHandler) Handle(req *codec.Packet) (*codec.Packet, error) {
	switch req.Cmd {
	case "log":
		logger.Info("Got your request: ", req.ID)
		err := h.store.Set(req.ID.String(), req.Body)
		if err != nil {
			return nil, err
		}
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
