package logger

import (
	"edoex/environment/flags"
	"fmt"
	"log"
)

func changeConsoleColor(level LoggerLevel) {
	if !flags.NoColor {
		fmt.Printf("%s%s%s", prefix, level, suffix)
	}
}

func Log(logging ...any) {
	if len(logging) > 0 {
		log.Println(logging...)
	} else {
		fmt.Println()
	}
}

func Logf(format string, logging ...any) {
	Log(fmt.Sprintf(format, logging...))
}
