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
)

func WritePendulumEffect(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" || !card.HasSubType("pendulum") {
		return nil
	}

	fitsBoxHeight := false
	var fontSizeOffset float64 = 0

	boxWidth := BuildPositions.PendulumTextBox.Dx()
	boxHeight := BuildPositions.PendulumTextBox.Dy()

	var lineImages []image.Image
	var lineHeight int

	for !fitsBoxHeight {
		currentFontSize := defaultTextFontSize + fontSizeOffset

		lineHeight = int(math.Ceil(currentFontSize))
		lineImages = make([]image.Image, 0)

		face, err := imagesutils.GetFontFace(
			embedfiles.FontCardEffect, currentFontSize)
		if err != nil {
			return err
		}

		c := gg.NewContext(boxWidth, boxHeight)
		c.SetFontFace(face)

		lines := c.WordWrap(card.PendulumDescription, float64(boxWidth))
		linesBoxHeight := 0

		for _, line := range lines {
			lineImg := imagesutils.TransparentBackgroundText(
				line, color.Black, face, boxWidth, lineHeight)

			lineImages = append(lineImages, lineImg)
			linesBoxHeight += lineHeight
		}

		fitsBoxHeight = linesBoxHeight <= boxHeight
		fontSizeOffset -= textFontSizeDecrement
	}

	offset := 0
	for _, ln := range lineImages {
		linePosition := BuildPositions.PendulumTextBox.Min.Add(image.Pt(0, offset))
		imagesutils.DrawAt(img, ln, linePosition)
		offset += lineHeight
	}

	return nil
}
