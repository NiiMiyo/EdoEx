package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/sliceutils"
	"edoex/utils/stringsutils"
	"image"
	"image/color"
	"image/draw"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

var (
	abilitiesFontSize = float64(30)
	abilitiesRect     = image.Rect(77, 793, 662, 830)
)

func WriteMonsterAbilities(img draw.Image, card *models.Card) (draw.Image, error) {
	if card.CardType != "monster" {
		return img, nil
	}

	fontFace, err := imagesutils.GetFontFace(
		embedfiles.FontCardMonsterDescription, abilitiesFontSize)
	if err != nil {
		return nil, err
	}

	context := gg.NewContextForImage(img)
	context.SetColor(color.Black)
	context.SetFontFace(fontFace)

	str := getAbilitiesString(card)

	w, h := context.MeasureString(str)
	abilitiesImg := imagesutils.TransparentBackgroundText(
		str, color.Black, fontFace, int(w), int(h))

	if w > float64(abilitiesRect.Dx()) {
		abilitiesImg = resize.Resize(
			uint(abilitiesRect.Dx()), uint(h), abilitiesImg, resize.Bilinear)
	}

	draw.Draw(img, abilitiesRect, abilitiesImg, abilitiesImg.Bounds().Min, draw.Over)

	return img, nil
}

func getAbilitiesString(card *models.Card) string {
	const sep = "/"

	toSeparate := sliceutils.Concatenate(card.SubTypes, card.Race)

	toSeparate = sliceutils.Filter(toSeparate, func(a string) bool {
		return !sliceutils.Contains(hiddenAbilities, a)
	})

	toSeparate = sliceutils.Map(toSeparate, func(s string) string {
		s = strings.ReplaceAll(s, "_", " ")

		sliced := strings.Split(s, " ")
		for i, v := range sliced {
			sliced[i] = stringsutils.Capitalize(v)
		}

		return strings.Join(sliced, " ")
	})

	return "[" + strings.Join(toSeparate, sep) + "]"
}

var hiddenAbilities = []string{"special_summon"}
