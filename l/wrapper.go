package l

func NewWrapper(l Logger) Wrapper {
	return &wrapperImpl{
		l: l,
	}
}

type wrapperImpl struct {
	l Logger
}

func (wrapper *wrapperImpl) GetLogger() Logger {
	return wrapper.l
}

func (wrapper *wrapperImpl) WithFields(fields ...Field) Wrapper {
	l := wrapper.l.WithFields(fields...)

	return NewWrapper(l)
}

func (wrapper *wrapperImpl) Fatal(a ...interface{}) {
	wrapper.l.Log(LevelFatal, a...)
}

func (wrapper *wrapperImpl) Fatalf(format string, a ...interface{}) {
	wrapper.l.Logf(LevelFatal, format, a...)
}

func (wrapper *wrapperImpl) Error(a ...interface{}) {
	wrapper.l.Log(LevelError, a...)
}

func (wrapper *wrapperImpl) Errorf(format string, a ...interface{}) {
	wrapper.l.Logf(LevelError, format, a...)
}

func (wrapper *wrapperImpl) Warn(a ...interface{}) {
	wrapper.l.Log(LevelWarn, a...)
}

func (wrapper *wrapperImpl) Warnf(format string, a ...interface{}) {
	wrapper.l.Logf(LevelWarn, format, a...)
}

func (wrapper *wrapperImpl) Info(a ...interface{}) {
	wrapper.l.Log(LevelInfo, a...)
}

func (wrapper *wrapperImpl) Infof(format string, a ...interface{}) {
	wrapper.l.Logf(LevelInfo, format, a...)
}

func (wrapper *wrapperImpl) Debug(a ...interface{}) {
	wrapper.l.Log(LevelDebug, a...)
}

func (wrapper *wrapperImpl) Debugf(format string, a ...interface{}) {
	wrapper.l.Logf(LevelDebug, format, a...)
}
