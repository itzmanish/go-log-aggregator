package handler

import (
	"github.com/itzmanish/go-log-aggregator/internal/codec"
	"github.com/itzmanish/go-log-aggregator/internal/logger"
)

type testHandler struct {
}

func NewTHandler() *testHandler {
	return &testHandler{}
}

func (th *testHandler) Handle(req *codec.Packet) (*codec.Packet, error) {
	logger.Info(req)
	return req, nil
}
