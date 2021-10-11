package logger

type Level int

const (
	LevelError = 1
	LevelWarn  = 2
	LevelInfo  = 3
	LevelDebug = 4
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
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Warn(a ...interface{})
	Warnf(format string, a ...interface{})
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
}
