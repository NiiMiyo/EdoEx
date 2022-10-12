package logger

import "fmt"

func Warn(logging ...any) {
	changeConsoleColor(WarningLevel)
	Log(logging...)
	changeConsoleColor(resetColor)
}

func Warnf(format string, logging ...any) {
	Warn(fmt.Sprintf(format, logging...))
}
