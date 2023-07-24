package logger

type LoggerLevel string

const (
	InfoLevel    LoggerLevel = "94"
	WarningLevel LoggerLevel = "33"
	ErrorLevel   LoggerLevel = "31"
	VerboseLevel LoggerLevel = "90"
	resetColor   LoggerLevel = ""

	prefix string = "\033["
	suffix string = "m"
)
