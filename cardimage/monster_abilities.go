package cardimage

import (
	"edoex/embedfiles"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/sliceutils"
	"edoex/utils/stringsutils"
	"image/color"
	"image/draw"
	"strings"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

const abilitiesFontSize float64 = 25.5

func WriteMonsterAbilities(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" {
		return nil
	}

	fontFace, err := imagesutils.GetFontFace(
		embedfiles.FontCardMonsterDescription, abilitiesFontSize)
	if err != nil {
		return err
	}

	context := gg.NewContext(0, 0)
	context.SetFontFace(fontFace)

	str := getAbilitiesString(card)

	abilitiesImg := imagesutils.TransparentBackgroundText(
		str, color.Black, fontFace)
	w := abilitiesImg.Bounds().Dx()

	if maxW := BuildPositions.AbilitiesBox.Dx(); w > maxW {
		abilitiesImg = resize.Resize(
			uint(maxW), uint(abilitiesImg.Bounds().Dy()), abilitiesImg, resize.Bilinear)
	}

	imagesutils.DrawAt(img, abilitiesImg, BuildPositions.AbilitiesBox.Min)

	return nil
}

func getAbilitiesString(card *models.Card) string {
	const sep = " / "

	toSeparate := sliceutils.Concatenate(card.Race, card.SubTypes)

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
