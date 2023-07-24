package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutHologram(img draw.Image, card *models.Card) error {
	logger.Verbosef("%d - Putting hologram", card.Id)
	hologramImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.GlobalTemplatesPath(), "hologram.png"))

	if err != nil {
		return err
	}

	imagesutils.DrawAt(img, hologramImage, BuildPositions.Hologram)

	return nil
}
