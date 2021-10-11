package logger

import "strings"

type Recorder interface {
	Log(level Level, a ...interface{})
	Logf(level Level, format string, a ...interface{})
}

type CommLogger struct {
	level    Level
	fields   []Field
	recorder Recorder
}

func NewCommLogger(recorder Recorder) Logger {
	return &CommLogger{
		level:    LevelInfo,
		fields:   nil,
		recorder: recorder,
	}
}

func (l *CommLogger) clone() *CommLogger {
	newLogger := &CommLogger{}
	newLogger.level = l.level
	newLogger.fields = append(newLogger.fields, l.fields...)
	newLogger.recorder = l.recorder
	return newLogger
}

func (l *CommLogger) SetLevel(level Level) {
	l.level = level
}

func (l *CommLogger) WithFields(fields ...Field) Logger {
	if len(fields) == 0 {
		return l
	}

	newLogger := l.clone()
	newLogger.fields = append(newLogger.fields, fields...)

	return newLogger
}

func (l *CommLogger) levelShouldRecord(level Level) bool {
	return level <= l.level
}

func (l *CommLogger) mapFields(fields ...Field) string {
	ss := &strings.Builder{}
	ss.WriteString("[")
	var f bool
	for _, field := range fields {
		switch field.T {
		case FieldTypeString:
			if f {
				ss.WriteString(" ")
			}
			f = true
			ss.WriteString(field.K)
			ss.WriteString(":")
			ss.WriteString(field.V.(string))
		case FieldTypeError:
			if f {
				ss.WriteString(" ")
			}
			f = true
			ss.WriteString(field.K)
			ss.WriteString(":")
			ss.WriteString(field.V.(error).Error())
		}
	}
	ss.WriteString("]")

	return ss.String()
}

func (l *CommLogger) Log(level Level, a ...interface{}) {
	if !l.levelShouldRecord(level) {
		return
	}
	l.recorder.Log(level, append([]interface{}{l.mapFields(l.fields...)}, a...)...)
}

func (l *CommLogger) Logf(level Level, format string, a ...interface{}) {
	if !l.levelShouldRecord(level) {
		return
	}
	l.recorder.Logf(level, l.mapFields(l.fields...)+" "+format, a...)
}
