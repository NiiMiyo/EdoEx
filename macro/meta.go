package macro

import (
	"edoex/environment"
	"edoex/models"
	"strconv"
	"strings"
)

var MetaMacro EdoexMacro = EdoexMacro{
	Name: "meta",
	Action: func(_ *models.Card, params []string) string {
		mentions := []string{}

		for _, p := range params {
			mentionedMeta := environment.MetasAlias[p]
			if mentionedMeta != nil {
				mentions = append(mentions, mentionedMeta.Name)
				continue
			}

			code, err := strconv.ParseInt(p, 0, 64)
			if err != nil {
				mentions = append(mentions, p)
				continue
			}

			mentionedMeta = environment.MetasIds[code]
			if mentionedMeta != nil {
				mentions = append(mentions, mentionedMeta.Name)
			}
		}

		return strings.Join(mentions, ", ")
	},
}
