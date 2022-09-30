package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

const (
	defaultTextFontSize           = 20
	textFontSizeDecrement float64 = .1
)

func WriteCardText(img draw.Image, card *models.Card) error {
	textBox := getCardTextBox(card)
	boxWidth := textBox.Dx()
	boxHeight := textBox.Dy()
	fitsBoxHeight := false
	var fontSizeOffset float64 = 0

	var lineImages []image.Image
	var lineHeight int

	for !fitsBoxHeight {
		currentFontSize := defaultTextFontSize + fontSizeOffset

		lineHeight = int(math.Ceil(currentFontSize))
		lineImages = make([]image.Image, 0)

		face, err := getCardTextFontFace(card, fontSizeOffset)
		if err != nil {
			return err
		}

		c := gg.NewContext(boxWidth, boxHeight)
		c.SetFontFace(face)

		lines := c.WordWrap(card.Description, float64(boxWidth))
		linesBoxHeight := 0

		for _, line := range lines {
			lineImg := imagesutils.TransparentBackgroundText(line, color.Black, face)

			lineImages = append(lineImages, lineImg)
			linesBoxHeight += lineHeight
		}

		fitsBoxHeight = linesBoxHeight <= boxHeight
		fontSizeOffset -= textFontSizeDecrement
	}

	offset := 0
	for _, ln := range lineImages {
		linePosition := textBox.Min.Add(image.Pt(0, offset))
		imagesutils.DrawAt(img, ln, linePosition)
		offset += lineHeight
	}

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
