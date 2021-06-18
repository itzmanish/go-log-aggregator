package logger

type Logger interface {
	Init(opts ...Option) error
	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Trace(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Errorf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	String() string
}

var DefaultLogger Logger = newLogger(WithLevel(DebugLevel))

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Log(level Level, args ...interface{}) {
	DefaultLogger.Log(level, args...)
}
func Logf(level Level, format string, args ...interface{}) {
	DefaultLogger.Logf(level, format, args...)
}
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}
func Trace(args ...interface{}) {
	DefaultLogger.Trace(args...)
}
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}
func Tracef(format string, args ...interface{}) {
	DefaultLogger.Tracef(format, args...)
}
func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}
func Panicf(format string, args ...interface{}) {
	DefaultLogger.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatalf(format, args...)
}
func String() string {
	return DefaultLogger.String()
}
