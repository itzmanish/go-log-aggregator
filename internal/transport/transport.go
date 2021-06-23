package transport

import (
	"encoding/json"
	"time"
)

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
	ID string `json:"id"`
	// Cmd is command containing the packets
	Cmd string `json:"cmd"`
	// Body is body of associated with command
	Body *LogBody `json:"body"`
	// Ack acknowledges for sent request
	Ack bool
	// Timestamp is time of packet creation.
	Timestamp time.Time `json:"timestamp"`
}

func (p *Packet) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
