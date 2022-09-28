package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/stringsutils"
	"fmt"
	"image/color"
	"image/draw"

	"github.com/fogleman/gg"
)

func PutCode(img draw.Image, card *models.Card) error {
	code := stringsutils.LeftJustify(fmt.Sprint(card.Id), 8, '0')

	face, err := imagesutils.GetFontFace(embedfiles.FontCardName, 30)
	if err != nil {
		return err
	}

	c := gg.NewContext(0, 0)
	c.SetFontFace(face)
	w, h := c.MeasureString(code)
	codeImg := imagesutils.TransparentBackgroundText(
		code, color.Black, face, int(w), int(h))

	imagesutils.DrawAt(img, codeImg, BuildPositions.Code)
	return nil
}
