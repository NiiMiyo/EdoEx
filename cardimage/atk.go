package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/stringsutils"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
)

const (
	atkFontSize float64 = defFontSize
	atkGap      int     = 15
)

func PutATK(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" {
		return nil
	}

	atkImage, err := getAtkImage(card)
	if err != nil {
		return err
	}

	atkWidth := atkImage.Bounds().Dx()

	atkPosition := BuildPositions.Defense.Sub(
		image.Pt(defLinkRatingOffset+atkGap+atkWidth, 0))

	imagesutils.DrawAt(img, atkImage, atkPosition)
	return nil
}

func getAtkImage(card *models.Card) (image.Image, error) {
	face, err := imagesutils.GetFontFace(embedfiles.FontCardAtkDef, atkFontSize)
	if err != nil {
		return nil, err
	}

	var text string
	if card.Atk >= 0 {
		text = "ATK/ " + stringsutils.LeftJustify(fmt.Sprint(card.Atk), 4, ' ')
	} else {
		text = "ATK/    ?"
	}

	c := gg.NewContext(0, 0)
	c.SetFontFace(face)
	w, h := c.MeasureString(text)

	return imagesutils.TransparentBackgroundText(
		text, color.Black, face, int(w), int(h)), nil
}
