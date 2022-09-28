package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
)

const pendulumScaleFontSize float64 = 37.5

func PutPendulumScale(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" {
		return nil
	}
	if !card.HasSubType("pendulum") {
		return nil
	}

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

	c := gg.NewContext(0, 0)
	c.SetFontFace(face)
	w, h := c.MeasureString(fmt.Sprint(scale))

	return imagesutils.TransparentBackgroundText(
		fmt.Sprint(scale), color.Black, face, int(w), int(h)), nil
}
