package cardimage

import (
	"edoex/embedfiles"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/stringsutils"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

const (
	defFontSize        float64 = 43
	linkRatingFontSize float64 = 30.5
)

func PutDefOrLinkRating(img draw.Image, card *models.Card) error {
	if card.CardType != "monster" {
		return nil
	}

	defLinkRatingImage, defLinkRatingPosition, err := getDefLinkRatingImage(card)
	if err != nil {
		return err
	}

	defLinkRatingOffset := defLinkRatingImage.Bounds().Dx()
	defLinkPos := defLinkRatingPosition.Sub(image.Pt(defLinkRatingOffset, 0))

	imagesutils.DrawAt(img, defLinkRatingImage, defLinkPos)
	return nil
}

func getDefLinkRatingImage(card *models.Card) (image.Image, image.Point, error) {
	var text string
	var fontBytes []byte
	var fontSize float64
	var defLinkRatingPosition image.Point

	if card.HasSubType("link") {
		logger.Verbosef("%d - Putting Link rating", card.Id)
		text = "LINK-" + fmt.Sprint(card.Level)
		fontBytes = embedfiles.FontCardLinkRating
		fontSize = linkRatingFontSize
		defLinkRatingPosition = BuildPositions.LinkRating
	} else {
		logger.Verbosef("%d - Putting Defense", card.Id)
		if card.Def >= 0 {
			text = "DEF/ " + stringsutils.LeftJustify(fmt.Sprint(card.Def), 4, ' ')
		} else {
			text = "DEF/    ?"
		}

		fontBytes = embedfiles.FontCardAtkDef
		fontSize = defFontSize
		defLinkRatingPosition = BuildPositions.Defense
	}

	face, err := imagesutils.GetFontFace(fontBytes, fontSize)
	if err != nil {
		return nil, defLinkRatingPosition, err
	}

	return imagesutils.TransparentBackgroundText(text, color.Black, face),
		defLinkRatingPosition, nil
}
