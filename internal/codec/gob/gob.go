package gob

import (
	"encoding/gob"
	"io"

	"github.com/itzmanish/go-log-aggregator/internal/codec"
)

type gobCodec struct {
	conn    io.ReadWriteCloser
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewGobCodec() codec.Codec {
	return &gobCodec{}
}

func (bc *gobCodec) Init(conn io.ReadWriteCloser) {
	bc.conn = conn
	bc.encoder = gob.NewEncoder(conn)
	bc.decoder = gob.NewDecoder(conn)
}

func (bc *gobCodec) Read(to interface{}) error {
	return bc.decoder.Decode(to)
}

func (bc *gobCodec) Write(to interface{}) error {
	return bc.encoder.Encode(to)
}

func (bc *gobCodec) Close() error {
	return bc.conn.Close()
}

func (bc *gobCodec) String() string {
	return "gob"
}
