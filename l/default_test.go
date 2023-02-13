package l

import "testing"

func TestConsoleLoggerWrapper(t *testing.T) {
	logger := NewConsoleLoggerWrapper()
	logger.GetLogger().SetLevel(LevelWarn)

	logger.Debug("a")
	logger.Info("b")
	logger.Warn("c")
	logger.Error("d")
}

func TestFileLoggerWrapper(t *testing.T) {
	logger := NewFileLoggerWrapper("logs/file.log")
	logger.GetLogger().SetLevel(LevelWarn)

	logger.Debug("a")
	logger.Info("b")
	logger.Warn("c")
	logger.Error("d")
}

func TestFileLoggerWrapper2(t *testing.T) {
	logger := NewCommLoggerEx(true, NewFileRecorder("logs/file.log"))
	logger.SetLevel(LevelInfo)

	lw := NewWrapper(logger)
	lw.Debug("a")
	lw.Info("b")
	lw.Warn("c")
	lw.Error("d")
	lw.Errorf("e:%d", 43)
}
