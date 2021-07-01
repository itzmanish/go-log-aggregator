package handler

import "github.com/itzmanish/go-log-aggregator/internal/codec"

// Handler handles request from client
type Handler interface {
	Handle(in *codec.Packet) (*codec.Packet, error)
}

type Options struct{}

type Option func(*Options)
