package transport

import (
	"encoding/json"
	"time"
)

// Packet is type of every packet that gets transported from client to server to client.
type Packet struct {
	// ID of packet
	ID string `json:"id"`
	// Cmd is command containing the packets
	Cmd string `json:"cmd"`
	// Body is body of associated with command
	Body interface{} `json:"body"`
	// Timestamp is time of packet creation.
	Timestamp time.Time `json:"timestamp"`
}

func (p *Packet) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
