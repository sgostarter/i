package l

import "github.com/sgostarter/i/logger"

func NewNopLoggerWrapper() logger.Wrapper {
	return logger.NewNopLoggerWrapper()
}

func NewConsoleLoggerWrapper() logger.Wrapper {
	return logger.NewConsoleLoggerWrapper()
}

func NewWrapper(l logger.Logger) logger.Wrapper {
	return logger.NewWrapper(l)
}

func NewWrapperWithLevel(l logger.Logger, level logger.Level) logger.Wrapper {
	if l == nil {
		return nil
	}

	l.SetLevel(level)

	return logger.NewWrapper(l)
}
