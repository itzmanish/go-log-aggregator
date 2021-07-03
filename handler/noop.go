package handler

import (
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/logger"
)

type noopHandler struct {
}

func NewTHandler() *noopHandler {
	return &noopHandler{}
}

func (th *noopHandler) Handle(req *codec.Packet) (*codec.Packet, error) {
	logger.Info(req)
	return req, nil
}
