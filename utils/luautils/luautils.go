package luautils

import (
	"regexp"
	"strings"

	"edoex/logger"
	"edoex/utils/sliceutils"
)

func MinifyCode(code []byte) []byte {
	codeLines := strings.Split(string(code), "\n")

	for _, re := range exprReplace {
		logger.Verbosef("Minifying Lua: %s", re.name)

		for ln, line := range codeLines {
			for {
				loc := re.expression.FindStringIndex(line)
				if loc == nil {
					break
				}

				line = line[:loc[0]] + re.replaceWith + line[loc[1]:]
			}

			codeLines[ln] = line
		}
	}

	return []byte(strings.Join(sliceutils.Filter(codeLines, lineFilter), "\n"))
}

func lineFilter(ln string) bool {
	return len(ln) > 0
}

type regexReplace struct {
	name        string
	expression  *regexp.Regexp
	replaceWith string
}

var exprReplace = []regexReplace{
	{
		"Removing indentation",
		regexp.MustCompile("^\\s+"),
		"",
	},
	{
		"Removing double spaces",
		regexp.MustCompile("\\s\\s"),
		"",
	},
	{
		"Removing comments",
		regexp.MustCompile("^--.*"),
		"",
	},
	{
		"Removing trailing whitespace",
		regexp.MustCompile("\\s+$"),
		"",
	},
	{
		"Removing whitespace before ','",
		regexp.MustCompile("\\s+,"),
		",",
	},
	{
		"Removing whitespace after ','",
		regexp.MustCompile(",\\s+"),
		",",
	},
	{
		"Removing whitespace before ';'",
		regexp.MustCompile("\\s+;"),
		";",
	},
	{
		"Removing whitespace after ';'",
		regexp.MustCompile(";\\s+"),
		";",
	},
	{
		"Removing whitespace before '('",
		regexp.MustCompile("\\s+\\("),
		"(",
	},
	{
		"Removing whitespace after '('",
		regexp.MustCompile("\\(\\s+"),
		"(",
	},
	{
		"Removing whitespace before ')'",
		regexp.MustCompile("\\s+\\)"),
		")",
	},
	{
		"Removing whitespace after ')'",
		regexp.MustCompile("\\)\\s+"),
		")",
	},
	{
		"Removing whitespace before '.'",
		regexp.MustCompile("\\s+\\."),
		".",
	},
	{
		"Removing whitespace after '.'",
		regexp.MustCompile("\\.\\s+"),
		".",
	},
	{
		"Removing whitespace before ','",
		regexp.MustCompile("\\s+\\,"),
		",",
	},
	{
		"Removing whitespace after ','",
		regexp.MustCompile("\\,\\s+"),
		",",
	},
	{
		"Removing whitespace before '+'",
		regexp.MustCompile("\\s+\\+"),
		"+",
	},
	{
		"Removing whitespace after '+'",
		regexp.MustCompile("\\+\\s+"),
		"+",
	},
	{
		"Removing whitespace before '='",
		regexp.MustCompile("\\s+\\="),
		"=",
	},
	{
		"Removing whitespace after '='",
		regexp.MustCompile("\\=\\s+"),
		"=",
	},
	{
		"Removing whitespace before '{'",
		regexp.MustCompile("\\s+\\{"),
		"{",
	},
	{
		"Removing whitespace after '{'",
		regexp.MustCompile("\\{\\s+"),
		"{",
	},
	{
		"Removing whitespace before '}'",
		regexp.MustCompile("\\s+\\}"),
		"}",
	},
	{
		"Removing whitespace after '}'",
		regexp.MustCompile("\\}\\s+"),
		"}",
	},
	{
		"Removing whitespace before '-'",
		regexp.MustCompile("\\s+\\-"),
		"-",
	},
	{
		"Removing whitespace after '-'",
		regexp.MustCompile("\\-\\s+"),
		"-",
	},
	{
		"Removing whitespace before '/'",
		regexp.MustCompile("\\s+\\/"),
		"/",
	},
	{
		"Removing whitespace after '/'",
		regexp.MustCompile("\\/\\s+"),
		"/",
	},
	{
		"Removing whitespace before '*'",
		regexp.MustCompile("\\s+\\*"),
		"*",
	},
	{
		"Removing whitespace after '*'",
		regexp.MustCompile("\\*\\s+"),
		"*",
	},
	{
		"Removing whitespace before '<'",
		regexp.MustCompile("\\s+\\<"),
		"<",
	},
	{
		"Removing whitespace after '<'",
		regexp.MustCompile("\\<\\s+"),
		"<",
	},
	{
		"Removing whitespace before '>'",
		regexp.MustCompile("\\s+\\>"),
		">",
	},
	{
		"Removing whitespace after '>'",
		regexp.MustCompile("\\>\\s+"),
		">",
	},
	{
		"Removing whitespace before ':'",
		regexp.MustCompile("\\s+\\:"),
		":",
	},
	{
		"Removing whitespace after ':'",
		regexp.MustCompile("\\:\\s+"),
		":",
	},
}
