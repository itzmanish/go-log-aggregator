package binary

import (
	"encoding/binary"
	"io"

	"github.com/itzmanish/go-log-aggregator/internal/codec"
)

type binaryCodec struct {
	conn io.ReadWriteCloser
}

func NewBinaryCodec() codec.Codec {
	return &binaryCodec{}
}

func (bc *binaryCodec) Init(conn io.ReadWriteCloser) {
	bc.conn = conn
}

func (bc *binaryCodec) Read(to interface{}) error {
	return binary.Read(bc.conn, binary.LittleEndian, to)
}

func (bc *binaryCodec) Write(to interface{}) error {
	return binary.Write(bc.conn, binary.LittleEndian, to)
}

func (bc *binaryCodec) Close() error {
	return bc.conn.Close()
}

func (bc *binaryCodec) String() string {
	return "Binary codec"
}
