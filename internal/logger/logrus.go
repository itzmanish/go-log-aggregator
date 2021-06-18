package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	Opts   LogOption
	logrus *logrus.Logger
}

func (l *logrusLogger) Log(level Level, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Log(logrus.Level(level), args...)
}

func (l *logrusLogger) Logf(level Level, format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Logf(logrus.Level(level), format, args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Info(args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Debug(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Error(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Warn(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Panic(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Fatal(args...)
}

func (l *logrusLogger) Trace(args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Trace(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Infof(format, args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Debugf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Errorf(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Warnf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Panicf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Fatalf(format, args...)
}

func (l *logrusLogger) Tracef(format string, args ...interface{}) {
	entry := l.logrus.WithFields(l.Opts.Fields)
	entry.Data["file"] = fileInfo(3)
	entry.Tracef(format, args...)
}

func (l *logrusLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.Opts)
	}
	return l.Load()
}

func (l *logrusLogger) Load() error {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	l.logrus.SetFormatter(customFormatter)
	if l.Opts.Output != nil {
		l.logrus.SetOutput(l.Opts.Output)
	}
	l.logrus.SetLevel(logrus.Level(l.Opts.LogLevel))
	return nil
}

func (l *logrusLogger) String() string {
	return "Logrus logger"
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func newLogger(opts ...Option) Logger {
	log := logrus.New()
	l := &logrusLogger{
		logrus: log,
	}
	l.Init(opts...)
	return l
}
