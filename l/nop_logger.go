package l

import "fmt"

type NopLogger struct{}

func (l *NopLogger) SetLevel(_ Level) {

}

func (l *NopLogger) WithFields(_ ...Field) Logger {
	return l
}

func (l *NopLogger) Log(level Level, a ...interface{}) {
	if level == LevelFatal {
		panic(fmt.Sprint(a...))
	}
}

func (l *NopLogger) Logf(level Level, format string, a ...interface{}) {
	if level == LevelFatal {
		panic(fmt.Sprintf(format, a...))
	}
}
