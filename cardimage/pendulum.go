package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutPendulum(img draw.Image, card *models.Card) error {
	if !card.HasSubType("pendulum") {
		return nil
	}

	logger.Verbosef("%d - Putting pendulum box", card.Id)
	pendulumImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.GlobalTemplatesPath(), "pendulum", "medium.png"))
	if err != nil {
		return err
	}

	imagesutils.DrawOver(img, pendulumImage)

	return nil
}
