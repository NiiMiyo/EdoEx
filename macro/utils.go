package macro

import (
	"edoex/logger"
	"edoex/models"
	"edoex/utils/sliceutils"
	"strings"
)

func applyMacrosOnCard(card *models.Card) error {
	logger.Verbosef("Applying macros on '%s' (%d)", card.Name, card.Id)
	cardMacroEnableSelf = false // Prevents running the same macro indefinitely
	logger.Verbose("Running macro on name")
	newName, err := applyMacrosOnText(card.Name, card)
	if err != nil {
		return err
	}
	card.Name = newName
	cardMacroEnableSelf = true

	logger.Verbose("Running macro on description")
	newDescription, err := applyMacrosOnText(card.Description, card)
	if err != nil {
		return err
	}
	card.Description = newDescription

	logger.Verbose("Running macro on pendulum description")
	newPendulumDescription, err := applyMacrosOnText(card.PendulumDescription, card)
	if err != nil {
		return err
	}
	card.PendulumDescription = newPendulumDescription

	for i := range card.Strings {
		logger.Verbosef("Running macro on string %d", i)
		newString, err := applyMacrosOnText(card.Strings[i], card)
		if err != nil {
			return err
		}
		card.Strings[i] = newString
	}

	return nil
}

func applyMacrosOnText(text string, referenceCard *models.Card) (string, error) {
	for {
		loc := macroCompiledRegex.FindStringIndex(text)
		if loc == nil {
			break
		}

		start, end := loc[0], loc[1]

		macroMatch := text[start:end]
		macroText := macroMatch[2 : len(macroMatch)-1] // remove "${" and "}"

		split := getMacroSplitText(macroText)

		macro := Macros[split[0]]
		params := split[1:]

		replaceWith := macro.Action(referenceCard, params)

		text = text[:start] + replaceWith + text[end:]
	}

	return text, nil
}

func getMacroSplitText(macroText string) []string {
	separated := strings.Split(macroText, MacroSeparator)

	return sliceutils.Map(separated, func(s string) string {
		return strings.TrimSpace(s)
	})
}
