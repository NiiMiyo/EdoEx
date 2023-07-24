package logger

import (
	"edoex/environment/flags"
	"fmt"
)

func Verbose(logging ...any) {
	if flags.Verbose {
		changeConsoleColor(VerboseLevel)
		Log(logging...)
		changeConsoleColor(resetColor)
	}
}

func Verbosef(format string, logging ...any) {
	Verbose(fmt.Sprintf(format, logging...))
}
