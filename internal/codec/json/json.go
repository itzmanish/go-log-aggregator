// Package json provides a json codec
package json

import (
	"encoding/json"
	"io"

	"github.com/itzmanish/go-loganalyzer/internal/codec"
)

type jsonCodec struct {
	Conn    io.ReadWriteCloser
	Encoder *json.Encoder
	Decoder *json.Decoder
}

func (c *jsonCodec) Init(conn io.ReadWriteCloser) {
	c.Conn = conn
	c.Encoder = json.NewEncoder(conn)
	c.Decoder = json.NewDecoder(conn)
}

func (c *jsonCodec) Read(b interface{}) error {
	if b == nil {
		return nil
	}
	return c.Decoder.Decode(b)
}

func (c *jsonCodec) Write(data interface{}) error {
	if data == nil {
		return nil
	}
	return c.Encoder.Encode(data)
}

func (c *jsonCodec) Close() error {
	return c.Conn.Close()
}

func (c *jsonCodec) String() string {
	return "json"
}

func NewCodec() codec.Codec {
	return &jsonCodec{}
}
