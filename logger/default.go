package logger

func NewNopLoggerWrapper() Wrapper {
	return NewWrapper(&NopLogger{})
}

func NewConsoleLoggerWrapper() Wrapper {
	log := &CommLogger{}
	log.AddRecorder(&FmtRecorder{})

	return NewWrapper(log)
}
