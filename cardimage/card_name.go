package cardimage

import (
	"edoex/embedfiles"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/color"
	"image/draw"

	"github.com/nfnt/resize"
)

const nameFontSize = float64(60)

func WriteCardName(img draw.Image, card *models.Card) error {
	logger.Verbosef("%d - Writing name", card.Id)
	fontFace, err := imagesutils.GetFontFace(
		embedfiles.FontCardName, nameFontSize)
	if err != nil {
		return err
	}

	color := getCardNameColor(card)
	nameImg := imagesutils.TransparentBackgroundText(card.Name, color, fontFace)
	w := nameImg.Bounds().Dx()

	if maxW := BuildPositions.NameBox.Dx(); w > maxW {
		nameImg = resize.Resize(
			uint(maxW), uint(nameImg.Bounds().Dy()), nameImg, resize.Bilinear)
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
