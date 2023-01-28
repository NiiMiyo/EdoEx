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

	getStarPosition := getPositionFunction(card, starImg.Bounds().Dx())

	amountStars := card.Level
	if amountStars < minStars {
		amountStars = minStars
	} else if amountStars > maxStars {
		amountStars = maxStars
	}

	for count := 0; count < int(amountStars); count++ {
		imagesutils.DrawAt(img, starImg, getStarPosition(count))
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

func getPositionFunction(card *models.Card, starOffset int) func(int) image.Point {
	if card.HasSubType("xyz") {
		return func(count int) image.Point {
			return BuildPositions.RankStars.Add(image.Point{starOffset * count, 0})
		}
	} else {
		return func(count int) image.Point {
			return BuildPositions.LevelStars.Sub(image.Point{starOffset * count, 0})
		}
	}
}
