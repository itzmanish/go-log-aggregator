package logger

import "io"

type Option func(o *LogOption)

type LogOption struct {
	// Out
	Output io.Writer
}

// WithOutput set default default output for the logger
func WithOutput(out io.Writer) Option {
	return func(o *LogOption) {
		o.Output = out
	}
}
