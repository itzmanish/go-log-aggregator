package logger

import "fmt"

type defaultLogger struct {
	Opts LogOption
}

func (l *defaultLogger) Log(args ...interface{}) {
	fmt.Fprintln(l.Opts.Output, args...)
}

func (l *defaultLogger) Logf(format string, args ...interface{}) {
	fmt.Fprintf(l.Opts.Output, format, args...)
}

func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.Opts)
	}
	return nil
}

func (l *defaultLogger) String() string {
	return "Default Logger"
}

func newLogger(opts ...Option) Logger {
	l := &defaultLogger{}
	l.Init(opts...)
	return l
}
