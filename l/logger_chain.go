package l

import (
	"sync"
	"time"
)

type LoggerChain interface {
	Logger

	AppendLogger(log Logger)
}

func NewLoggerChain() LoggerChain {
	return &loggerChainImpl{
		loggersMap: make(map[Logger]interface{}),
	}
}

type loggerChainImpl struct {
	sync.RWMutex
	loggers    []Logger
	loggersMap map[Logger]interface{}
}

func (impl *loggerChainImpl) AppendLogger(log Logger) {
	impl.Lock()
	defer impl.Unlock()

	if _, ok := impl.loggersMap[log]; ok {
		return
	}

	impl.loggersMap[log] = time.Now()
	impl.loggers = append(impl.loggers, log)
}

func (impl *loggerChainImpl) SetLevel(level Level) {
	impl.RLock()
	defer impl.RUnlock()

	for _, log := range impl.loggers {
		log.SetLevel(level)
	}
}

func (impl *loggerChainImpl) WithFields(fields ...Field) Logger {
	impl.RLock()
	defer impl.RUnlock()

	newImpl := NewLoggerChain()

	for _, log := range impl.loggers {
		newImpl.AppendLogger(log.WithFields(fields...))
	}

	return newImpl
}

func (impl *loggerChainImpl) Log(level Level, a ...interface{}) {
	impl.RLock()
	defer impl.RUnlock()

	for _, log := range impl.loggers {
		log.Log(level, a...)
	}
}

func (impl *loggerChainImpl) Logf(level Level, format string, a ...interface{}) {
	impl.RLock()
	defer impl.RUnlock()

	for _, log := range impl.loggers {
		log.Logf(level, format, a...)
	}
}
