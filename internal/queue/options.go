package queue

import (
	"time"

	"github.com/itzmanish/go-log-aggregator/internal/client"
)

// Options defines available option field for queue
type Options struct {
	Interval     time.Duration
	Client       client.Client
	MaxQueueSize int
}

// Option is a function which takes pointer of Options
type Option func(*Options)

// WithTimeInterval sets time interval for two flush execution
func WithTimeInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.Interval = interval
	}
}

// WithClient sets client for queue
func WithClient(client client.Client) Option {
	return func(o *Options) {
		o.Client = client
	}
}

// WithMaxQueueSize sets the max size of queue to hold
func WithMaxQueueSize(size int) Option {
	return func(o *Options) {
		o.MaxQueueSize = size
	}
}
