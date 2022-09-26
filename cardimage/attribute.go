package cardimage

import (
	"edoex/environment"
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

	attributePath = filepath.Join(environment.TemplatesPath(), "attributes", attributePath+".png")
	attributeImage, err := imagesutils.LoadImageFromPath(attributePath)
	if err != nil {
		return err
	}

	imagesutils.DrawAt(img, attributeImage, BuildPositions.attribute)
	return nil
}
