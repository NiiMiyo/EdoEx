package cardimage

import (
	"edoex/embedfiles"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

const pendulumScaleFontSize float64 = 37.5

func PutPendulumScale(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" || !card.HasSubType("pendulum") {
		return nil
	}

	logger.Verbosef("%d - Putting pendulum scales", card.Id)
	scaleImage, err := getPendulumScaleImage(card.Scale)
	if err != nil {
		return err
	}

	leftScalePosition := BuildPositions.PendulumScaleLeft.Sub(
		image.Pt(scaleImage.Bounds().Dx()/2, 0))

	rightScalePosition := BuildPositions.PendulumScaleRight.Sub(
		image.Pt(scaleImage.Bounds().Dx()/2, 0))

	imagesutils.DrawAt(img, scaleImage, leftScalePosition)
	imagesutils.DrawAt(img, scaleImage, rightScalePosition)
	return nil
}

func getPendulumScaleImage(scale int64) (image.Image, error) {
	face, err := imagesutils.GetFontFace(
		embedfiles.FontCardAtkDef, pendulumScaleFontSize)
	if err != nil {
		return nil, err
	}

	return imagesutils.TransparentBackgroundText(
		fmt.Sprint(scale), color.Black, face), nil
}
