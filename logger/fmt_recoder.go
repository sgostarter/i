package logger

import "fmt"

type FmtRecorder struct{}

func (r *FmtRecorder) Log(_ Level, a ...interface{}) {
	// nolint: forbidigo
	fmt.Println(a...)
}
func (r *FmtRecorder) Logf(_ Level, format string, a ...interface{}) {
	if len(format) > 0 {
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
	}

	// nolint: forbidigo
	fmt.Printf(format, a...)
}
