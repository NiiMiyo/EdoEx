package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"image/draw"
	"path/filepath"
)

const (
	minStars = 0
	maxStars = 12
)

func PutStars(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" || card.HasSubType("link") {
		return nil
	}

	starImg, err := getStarImage(card)
	if err != nil {
		return err
	}

	starOffset := image.Pt(starImg.Bounds().Dx(), 0)

	amountStars := card.Level
	if amountStars < minStars {
		amountStars = minStars
	} else if amountStars > maxStars {
		amountStars = maxStars
	}

	for count := 0; count < int(amountStars); count++ {
		starPosition := BuildPositions.Stars.Sub(starOffset.Mul(count))
		imagesutils.DrawAt(img, starImg, starPosition)
	}

	return nil
}

func getStarImage(card *models.Card) (image.Image, error) {
	var starFilename string
	if card.HasSubType("xyz") {
		starFilename = "rank.png"
	} else {
		starFilename = "level.png"
	}

	return imagesutils.LoadImageFromPath(
		filepath.Join(environment.GlobalTemplatesPath(), starFilename))
}
