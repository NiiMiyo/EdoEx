package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/sliceutils"
	"image/draw"
	"path/filepath"
)

func PutSpellTrapType(img draw.Image, card *models.Card) (draw.Image, error) {
	if card.CardType != "spell" && card.CardType != "trap" {
		return img, nil
	}

	specialTypes := []string{
		"continuous", "counter", "equip", "field", "quick_play", "ritual",
	}

	var typeToPut string
	for _, t := range card.SubTypes {
		if sliceutils.Contains(specialTypes, t) {
			typeToPut = t
			break
		}
	}

	typesDir := filepath.Join(environment.TemplatesPath(), "st_types")

	if typeToPut == "" { // normal spell/trap
		typePath := filepath.Join(
			typesDir, string(card.CardType[0])+"_normal.png",
		)

		typeImage, err := imagesutils.LoadImageFromPath(typePath)
		if err != nil {
			return nil, err
		}

		return imagesutils.DrawOver(img, typeImage), nil
	}

	iconPath := filepath.Join(typesDir, typeToPut+".png")
	textPath := filepath.Join(typesDir, string(card.CardType[0])+"_typed.png")

	iconImage, err := imagesutils.LoadImageFromPath(iconPath)
	if err != nil {
		return nil, err
	}
	img = imagesutils.DrawOver(img, iconImage)

	textImage, err := imagesutils.LoadImageFromPath(textPath)
	if err != nil {
		return nil, err
	}

	return imagesutils.DrawOver(img, textImage), nil
}
