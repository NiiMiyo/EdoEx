package logger

import "fmt"

func Error(logging ...any) {
	changeConsoleColor(ErrorLevel)
	Log(logging...)
	changeConsoleColor(resetColor)
}

func Errorf(format string, logging ...any) {
	Error(fmt.Sprintf(format, logging...))
}

func ErrorErr(message string, err error) {
	Error(fmt.Sprintf(message+":\n\t%s\n", err))
}

func ErrorfErr(message string, err error, logging ...any) {
	ErrorErr(fmt.Sprintf(message, logging...), err)
}
