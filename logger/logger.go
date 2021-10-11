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
