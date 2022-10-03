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
			var mentionedCard *models.Card
			if p == "self" {
				mentionedCard = card
			} else {
				code, err := strconv.ParseInt(p, 0, 64)
				if err == nil {
					mentionedCard = environment.Cards[code]
				}
			}

			if mentionedCard != nil && (mentionedCard.Id != card.Id || cardMacroEnableSelf) {
				mentions = append(mentions, mentionedCard.Name)
			} else {
				mentions = append(mentions, p)
			}
		}

		return strings.Join(mentions, ", ")
	},
}
