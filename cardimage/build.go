package cardimage

import (
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"image/draw"
)

var cropMargin = image.Rect(23, 24, 716, 1036)

type BuildImageFunction func(draw.Image, *models.Card) error

var buildFunctions = []BuildImageFunction{
	// ! Remember to put link arrows on the end
	PutAttribute, PutSpellTrapType, PutArtwork, PutPendulum, WriteCardName,
	WriteMonsterAbilities, PutDefOrLinkRating, PutStars, WriteCardText,
	PutLinkArrows,
}

func BuildCardImage(card *models.Card) (draw.Image, error) {
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

	return imagesutils.Crop(img, cropMargin), nil
}
