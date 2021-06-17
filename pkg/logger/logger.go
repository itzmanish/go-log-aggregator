package logger

import "os"

type Logger interface {
	Init(opts ...Option) error
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	String() string
}

var DefaultLogger Logger = newLogger(WithOutput(os.Stderr))

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Log(args ...interface{}) {
	DefaultLogger.Log(args...)
}

func Logf(format string, args ...interface{}) {
	DefaultLogger.Logf(format, args...)
}
