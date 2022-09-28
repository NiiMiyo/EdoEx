package cardimage

import (
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
)

type BuildImageFunction func(draw.Image, *models.Card) error
type BuildStaticImageFunction func(draw.Image) error

var buildFunctions = []BuildImageFunction{
	// ! Remember to put link arrows on the end
	PutAttribute, PutSpellTrapType, PutArtwork, PutPendulum, WriteCardName,
	WriteMonsterAbilities, PutDefOrLinkRating, PutATK, PutStars, WriteCardText,
	PutPendulumScale, WritePendulumEffect, PutCode, PutLinkArrows,
}

var buildStaticFunctions = []BuildStaticImageFunction{
	PutHologram, PutMadeWithEdoex,
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

	for _, f := range buildStaticFunctions {
		err = f(img)
		if err != nil {
			return nil, err
		}
	}

	return imagesutils.Crop(img, BuildPositions.CardImageBox), nil
}
