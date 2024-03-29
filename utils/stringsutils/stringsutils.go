package stringsutils

import "strings"

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func LeftJustify(s string, size int, completeWith rune) string {
	lenS := len(s)
	if lenS >= size {
		return s
	}

	return s + strings.Repeat(string(completeWith), size-lenS)
}
