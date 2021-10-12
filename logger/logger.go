package logger

type Level int

const (
	LevelFatal = 10
	LevelError = 11
	LevelWarn  = 12
	LevelInfo  = 13
	LevelDebug = 14
)

type Logger interface {
	SetLevel(level Level)
	WithFields(fields ...Field) Logger
	Log(level Level, a ...interface{})
	Logf(level Level, format string, a ...interface{})
}

type Wrapper interface {
	GetLogger() Logger

	WithFields(fields ...Field) Wrapper

	Fatal(a ...interface{})
	Fatalf(ormat string, a ...interface{})
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Warn(a ...interface{})
	Warnf(format string, a ...interface{})
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
}
