package handler

import (
	"encoding/json"
	"io"

	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/server"
	"github.com/itzmanish/go-loganalyzer/internal/store"
	"github.com/itzmanish/go-loganalyzer/internal/transport"
)

type srvHandler struct {
	store store.Store
}

func NewHandler(s store.Store) server.Handler {
	return &srvHandler{
		store: s,
	}
}

func (h *srvHandler) Handle(req *transport.Packet, w io.Writer) error {
	switch req.Cmd {
	case "log":
		logger.Info("Got your request: ", req.ID)
		err := h.store.Set(req.ID, req.Body)
		if err != nil {
			logger.Error(err)
			return err
		}
		ack := &transport.Packet{
			ID:  req.ID,
			Ack: true,
		}
		err = json.NewEncoder(w).Encode(ack)
		if err != nil {
			logger.Error(err)
			return err
		}

	default:
		logger.Info(req)
	}
	return nil
}
