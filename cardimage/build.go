package cardimage

import (
	"edoex/models"
	"image"
	"image/draw"
)

type BuildImageFunction func(draw.Image, *models.Card) error

var buildFunctions = []BuildImageFunction{
	// ! Remember to put link arrows on the end
	PutAttribute, PutSpellTrapType, PutArtwork, PutPendulum, WriteCardName,
	WriteMonsterAbilities, PutAtkDef, PutLinkArrows,
}

func BuildCardImage(card *models.Card) (image.Image, error) {
	img, err := GetCardBase(card)
	if err != nil {
		return nil, err
	}

	for _, f := range buildFunctions {
		err = f(img, card)
		if err != nil {
			return nil, err
		}
	}

	err = PutHologram(img)
	if err != nil {
		return nil, err
	}

	return CropMargin(img), nil
}