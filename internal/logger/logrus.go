package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	Opts   LogOption
	logrus *logrus.Logger
}

func (l *logrusLogger) Log(level Level, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Log(logrus.Level(level), args...)
}

func (l *logrusLogger) Logf(level Level, format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Logf(logrus.Level(level), format, args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Info(args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Debug(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Error(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Warn(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Panic(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Fatal(args...)
}

func (l *logrusLogger) Trace(args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Trace(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Infof(format, args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Debugf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Errorf(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Warnf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Panicf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Fatalf(format, args...)
}

func (l *logrusLogger) Tracef(format string, args ...interface{}) {
	l.logrus.WithFields(l.Opts.Fields).Tracef(format, args...)
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

func newLogger(opts ...Option) Logger {
	log := logrus.New()
	l := &logrusLogger{
		logrus: log,
	}
	l.Init(opts...)
	return l
}
