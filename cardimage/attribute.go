package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutAttribute(img draw.Image, card *models.Card) error {
	var attributePath string

	switch card.CardType {
	case "spell":
		attributePath = "spell"
	case "trap":
		attributePath = "trap"
	default:
		if len(card.Attribute) > 0 {
			attributePath = card.Attribute[0]
		}
	}

	if attributePath == "" {
		return nil
	}

	logger.Verbosef("%d - Putting attribute '%s'", card.Id, attributePath)
	attributePath = filepath.Join(environment.GlobalTemplatesPath(), "attributes", attributePath+".png")
	attributeImage, err := imagesutils.LoadImageFromPath(attributePath)
	if err != nil {
		return err
	}

	imagesutils.DrawAt(img, attributeImage, BuildPositions.Attribute)
	return nil
}
