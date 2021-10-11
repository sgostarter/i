package logger

import "fmt"

type FmtRecorder struct{}

func (r *FmtRecorder) Log(_ Level, a ...interface{}) {
	fmt.Println(a...)
}
func (r *FmtRecorder) Logf(_ Level, format string, a ...interface{}) {
	if len(format) > 0 {
		if format[len(format)-1] != '\n' {
			format += "\n"
		}
	}
	fmt.Printf(format, a...)
}
