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

var (
	atkDefFontSize = float64(43)
	atkDefRect     = image.Rect(390, 940, 658, 972)
)

func PutAtkDef(img draw.Image, card *models.Card) (draw.Image, error) {
	if card.CardType != "monster" {
		return img, nil
	}

	// Go doesn't support the correct font (values.ttf)
	// For now, this one will be used instead
	// TODO: Use correct font and readjust atkDefRect and atkDefSizeFont
	fontFace, err := imagesutils.GetFontFace(
		embedfiles.FontCardName, atkDefFontSize)
	if err != nil {
		return nil, err
	}

	context := gg.NewContextForImage(img)
	context.SetColor(color.Black)
	context.SetFontFace(fontFace)

	str := getAtkDefString(card)
	w, h := context.MeasureString(str)
	atkDefImg := imagesutils.TransparentBackgroundText(
		str, color.Black, fontFace, int(w), int(h))

	offset := (atkDefRect.Min.X + atkDefImg.Bounds().Dx()) - atkDefRect.Max.X
	var r image.Rectangle
	if offset > 0 {
		r = atkDefRect.Canon()
		r.Min.X -= offset
	} else {
		r = atkDefRect
	}

	draw.Draw(img, r, atkDefImg, atkDefImg.Bounds().Min, draw.Over)

	return img, nil
}

func getAtkDefString(card *models.Card) string {
	str := fmt.Sprintf("ATK/%s", getValueStr(card.Atk))

	if !card.HasSubType("link") {
		str += fmt.Sprintf(" DEF/%s", getValueStr(card.Def))
	}

	return str
}

func getValueStr(value int64) string {
	if value >= 0 {
		return stringsutils.LeftJustify(fmt.Sprint(value), 4, ' ')
	} else {
		return "   ?"
	}
}
