package server

import (
	"encoding/json"
	"io"
)

type Codec interface {
	Encode(data interface{}) error
	Decode(data []byte, res interface{}) error
	String() string
}

type jsonCodec struct {
	jEncoder *json.Encoder
	jDecoder *json.Decoder
}

func (jc *jsonCodec) Encode(data interface{}) error {
	return jc.jEncoder.Encode(data)
}

func (jc *jsonCodec) Decode(data []byte, res interface{}) error {
	return jc.jDecoder.Decode(res)
}

func (jc *jsonCodec) String() string {
	return "Json"
}

func NewCodec(src io.ReadWriter) Codec {
	j := &jsonCodec{
		jEncoder: json.NewEncoder(src),
		jDecoder: json.NewDecoder(src),
	}
	return j
}
