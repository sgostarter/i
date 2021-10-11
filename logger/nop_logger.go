package logger

type NopLogger struct{}

func (l *NopLogger) SetLevel(_ Level) {

}

func (l *NopLogger) WithFields(_ ...Field) Logger {
	return l
}

func (l *NopLogger) Log(_ Level, _ ...interface{}) {

}

func (l *NopLogger) Logf(_ Level, _ string, _ ...interface{}) {

}
