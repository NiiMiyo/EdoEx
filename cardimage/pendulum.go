package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"path/filepath"
)

func PutPendulum(img image.Image, card *models.Card) (image.Image, error) {
	if !card.HasSubType("pendulum") {
		return img, nil
	}

	pendulumImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.TemplatesPath(), "pendulum", "medium.png"))
	if err != nil {
		return nil, err
	}

	return imagesutils.DrawOver(img, pendulumImage), nil
}
