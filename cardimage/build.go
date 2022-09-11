package cardimage

import (
	"edoex/models"
	"image"
)

func BuildCardImage(card *models.Card) (image.Image, error) {
	img, err := GetCardBase(card)
	if err != nil {
		return nil, err
	}

	img, err = PutAttribute(img, card)
	if err != nil {
		return nil, err
	}

	return PutSpellTrapType(img, card)
}
