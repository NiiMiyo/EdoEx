package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

const nameFontSize = float64(60)

func WriteCardName(img draw.Image, card *models.Card) error {
	fontFace, err := imagesutils.GetFontFace(
		embedfiles.FontCardName, nameFontSize)
	if err != nil {
		return err
	}

	color := getCardNameColor(card)

	context := gg.NewContext(0, 0)
	context.SetFontFace(fontFace)

	w, h := context.MeasureString(card.Name)
	nameImg := imagesutils.TransparentBackgroundText(
		card.Name, color, fontFace, int(w), int(h))

	if maxW := float64(BuildPositions.NameBox.Dx()); w > maxW {
		nameImg = resize.Resize(uint(maxW), uint(h), nameImg, resize.Bilinear)
	}

	imagesutils.DrawAt(img, nameImg, BuildPositions.NameBox.Min)

	return nil
}

func getCardNameColor(card *models.Card) color.Color {
	if card.CardType == "spell" || card.CardType == "trap" {
		return color.White
	} else if card.HasSubType("xyz") || card.HasSubType("link") {
		return color.White
	}

	return color.Black
}
