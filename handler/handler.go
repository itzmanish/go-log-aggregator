package handler

import (
	"io"

	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/server"
	"github.com/itzmanish/go-loganalyzer/internal/transport"
)

type srvHandler struct{}

func NewHandler() server.Handler {
	return &srvHandler{}
}

func (h *srvHandler) Handle(req *transport.Packet, w io.Writer) error {
	logger.Info(req)
	return nil
}
