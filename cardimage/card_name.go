package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

var (
	nameDpi      = float64(72)
	nameFontSize = float64(60)
	nameRect     = image.Rect(76, 69, 592, 133)
)

func WriteCardName(img image.Image, card *models.Card) (image.Image, error) {
	fontFace, err := imagesutils.GetFontFace(
		embedfiles.FontCardName, nameFontSize, nameDpi)
	if err != nil {
		return nil, err
	}

	color := getCardNameColor(card)

	context := gg.NewContextForImage(img)
	context.SetColor(color)
	context.SetFontFace(fontFace)

	w, h := context.MeasureString(card.Name)
	nameImg := imagesutils.TransparentBackgroundText(
		card.Name, color, fontFace, int(w), int(h))

	if w > float64(nameRect.Dx()) {
		nameImg = resize.Resize(uint(nameRect.Dx()), uint(h), nameImg, resize.Bilinear)
	}

	drawable := imagesutils.GetRGBA(img)
	draw.Draw(drawable, nameRect, nameImg, nameImg.Bounds().Min, draw.Over)

	return drawable, nil
}

func getCardNameColor(card *models.Card) color.Color {
	if card.CardType == "spell" || card.CardType == "trap" {
		return color.White
	} else if card.HasSubType("xyz") || card.HasSubType("link") {
		return color.White
	}

	return color.Black
}
