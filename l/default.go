package l

func NewNopLoggerWrapper() Wrapper {
	return NewWrapper(&NopLogger{})
}

func NewConsoleLoggerWrapper() Wrapper {
	log := &CommLogger{
		level: LevelInfo,
	}
	log.AddRecorder(&FmtRecorder{})

	return NewWrapper(log)
}
