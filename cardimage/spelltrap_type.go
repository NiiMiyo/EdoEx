package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/sliceutils"
	"image/draw"
	"path/filepath"
)

func PutSpellTrapType(img draw.Image, card *models.Card) error {
	if card.CardType != "spell" && card.CardType != "trap" {
		return nil
	}

	specialTypes := []string{
		"continuous", "counter", "equip", "field", "quickplay", "ritual",
	}

	var typeToPut string
	for _, t := range card.SubTypes {
		if sliceutils.Contains(specialTypes, t) {
			typeToPut = t
			break
		}
	}

	logger.Verbosef("%d - Putting Spell/Trap type '%s'", card.Id, typeToPut)

	typesDir := filepath.Join(environment.GlobalTemplatesPath(), "st_types")

	if typeToPut == "" { // normal spell/trap
		typePath := filepath.Join(
			typesDir, string(card.CardType[0])+"_normal.png",
		)

		typeImage, err := imagesutils.LoadImageFromPath(typePath)
		if err != nil {
			return err
		}

		imagesutils.DrawAt(img, typeImage, BuildPositions.SpellTrapType)
		return nil
	}

	iconPath := filepath.Join(typesDir, typeToPut+".png")
	textPath := filepath.Join(typesDir, string(card.CardType[0])+"_typed.png")

	iconImage, err := imagesutils.LoadImageFromPath(iconPath)
	if err != nil {
		return err
	}
	imagesutils.DrawAt(img, iconImage, BuildPositions.SpellTrapType)

	textImage, err := imagesutils.LoadImageFromPath(textPath)
	if err != nil {
		return err
	}

	imagesutils.DrawAt(img, textImage, BuildPositions.SpellTrapType)

	return nil
}
