package macro

import (
	"edoex/environment"
	"edoex/models"
	"strconv"
	"strings"
)

var cardMacroEnableSelf bool = true

var CardMacro EdoexMacro = EdoexMacro{
	Name: "card",
	Action: func(card *models.Card, params []string) string {
		mentions := []string{}

		for _, p := range params {
			code, err := strconv.ParseInt(p, 0, 64)

			if err != nil {
				if p == "self" && cardMacroEnableSelf {
					mentions = append(mentions, card.Name)
				} else {
					mentions = append(mentions, p)
				}
				continue
			}

			mentionedCard := environment.Cards[code]
			if mentionedCard != nil {
				mentions = append(mentions, mentionedCard.Name)
			} else {
				mentions = append(mentions, p)
			}
		}

		return strings.Join(mentions, ", ")
	},
}
