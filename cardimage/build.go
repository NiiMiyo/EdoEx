package cardimage

import (
	"edoex/models"
	"image"
)

type BuildImageFunction func(image.Image, *models.Card) (image.Image, error)

var buildFunctions = []BuildImageFunction{
	// ! Remember to put link arrows on the end
	PutAttribute, PutSpellTrapType, PutArtwork, PutPendulum, PutLinkArrows,
}

func BuildCardImage(card *models.Card) (image.Image, error) {
	img, err := GetCardBase(card)
	if err != nil {
		return nil, err
	}

	for _, f := range buildFunctions {
		img, err = f(img, card)
		if err != nil {
			return nil, err
		}
	}

	return PutHologram(img)
}
