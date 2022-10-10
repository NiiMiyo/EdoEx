package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutPendulum(img draw.Image, card *models.Card) error {
	if !card.HasSubType("pendulum") {
		return nil
	}

	pendulumImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.GlobalTemplatesPath(), "pendulum", "medium.png"))
	if err != nil {
		return err
	}

	imagesutils.DrawOver(img, pendulumImage)

	return nil
}
