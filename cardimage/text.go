package cardimage

import (
	"edoex/embedfiles"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
)

const (
	defaultTextFontSize           = 20
	textFontSizeDecrement float64 = .1
)

func WriteCardText(img draw.Image, card *models.Card) error {
	logger.Verbosef("%d - Writing description box", card.Id)
	textBox := getCardTextBox(card)
	boxWidth := textBox.Dx()
	boxHeight := textBox.Dy()

	fitsBoxHeight := false
	var fontSizeOffset float64 = 0
	var textBoxImage image.Image

	for !fitsBoxHeight {
		face, err := getCardTextFontFace(card, fontSizeOffset)
		if err != nil {
			return err
		}

		textBoxImage = imagesutils.JustifiedText(card.Description, color.Black, face, boxWidth)
		fitsBoxHeight = textBoxImage.Bounds().Dy() <= boxHeight
		fontSizeOffset -= textFontSizeDecrement
	}

	imagesutils.DrawAt(img, textBoxImage, textBox.Min)
	return nil
}

func getCardTextBox(card *models.Card) image.Rectangle {
	if card.CardType == "spell" || card.CardType == "trap" {
		return BuildPositions.TextSpellTrapBox
	}

	return BuildPositions.TextMonsterBox
}

func getCardTextFontFace(card *models.Card, sizeOffset float64) (font.Face, error) {
	var fontBytes []byte

	if card.CardType == "spell" || card.CardType == "trap" ||
		card.HasAnySubType("effect", "xyz", "link", "synchro") {
		fontBytes = embedfiles.FontCardEffect
	} else {
		fontBytes = embedfiles.FontCardFlavorText
	}

	return imagesutils.GetFontFace(fontBytes, defaultTextFontSize+sizeOffset)
}
