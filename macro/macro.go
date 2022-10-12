package macro

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"fmt"
	"regexp"
	"strings"
)

var (
	macrosToBeApplied []EdoexMacro = []EdoexMacro{
		MetaMacro, CardMacro,
	}

	Macros             map[string]EdoexMacro = make(map[string]EdoexMacro)
	macroRegex         string
	macroCompiledRegex *regexp.Regexp
)

const MacroSeparator = ":"

type EdoexMacro struct {
	Name   string
	Action func(card *models.Card, params []string) string
}

func ApplyMacros() {
	nameList := make([]string, 0)

	for _, m := range macrosToBeApplied {
		Macros[m.Name] = m
		nameList = append(nameList, m.Name)
	}

	macroRegex = fmt.Sprintf(
		`\${ *(%s) *(?:\:[^}]*)* *}`,
		strings.Join(nameList, "|"),
	)

	var err error
	macroCompiledRegex, err = regexp.Compile(macroRegex)
	if err != nil {
		logger.ErrorErr("Error compiling macros RegExp. Macros will not be loaded", err)
		return
	}

	for _, c := range environment.Cards {
		err := applyMacrosOnCard(c)

		if err != nil {
			logger.ErrorfErr("Error running macros on '%d'", err, c.Id)
		}
	}
}
