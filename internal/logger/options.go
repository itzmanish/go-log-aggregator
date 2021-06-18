package logger

import (
	"io"
)

type Level uint32

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type Option func(o *LogOption)

type LogOption struct {
	// Out
	Output io.Writer
	// Fields
	Fields map[string]interface{}
	// Level
	LogLevel Level
}

// WithOutput set default default output for the logger
func WithOutput(out io.Writer) Option {
	return func(o *LogOption) {
		o.Output = out
	}
}

// WithFields set logger with custom fields
func WithFields(fields map[string]interface{}) Option {
	return func(o *LogOption) {
		o.Fields = fields
	}
}

// WithLevel sets logging level
func WithLevel(level Level) Option {
	return func(o *LogOption) {
		o.LogLevel = level
	}
}
