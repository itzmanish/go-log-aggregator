// Package codec is an interface for encoding messages
package codec

import (
	"errors"
	"io"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidMessage = errors.New("invalid message")
)

// Takes in a connection/buffer and returns a new Codec
type NewCodec func(io.ReadWriteCloser) Codec

type Codec interface {
	Reader
	Writer
	Init(io.ReadWriteCloser)
	Close() error
	String() string
}

type Reader interface {
	Read(interface{}) error
}

type Writer interface {
	Write(interface{}) error
}

// LogBody describe body of a single packets.
// It contains the name of log with log to store and tags
// along with timestamp field which contains the time of creation of log
type LogBody struct {
	Name      string                   `json:"name"`
	Log       string                   `json:"log"`
	Tags      []map[string]interface{} `json:"tags"`
	Timestamp time.Time                `json:"timestamp"`
}

// Packet is type of every packet that gets transported from client to server to client.
type Packet struct {
	// ID of packet
	ID uuid.UUID `json:"id"`
	// AgentID is id of agent from which logs are sent
	AgentID uuid.UUID `json:"agent_id"`
	// Cmd is command containing the packets
	Cmd string `json:"cmd"`
	// Body is body of associated with command
	Body *LogBody `json:"body"`
	// Ack acknowledges for sent request
	Ack bool
	// Error
	Error error
	// Timestamp is time of packet creation.
	Timestamp time.Time `json:"timestamp"`
}
