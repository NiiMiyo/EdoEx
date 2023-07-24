package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"fmt"
	"image"
	"image/draw"
	"path/filepath"

	"github.com/nfnt/resize"
)

func PutArtwork(img draw.Image, card *models.Card) error {
	logger.Verbosef("%d - Putting Artwork", card.Id)
	artworkImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.SourceArtworksPath(), fmt.Sprintf("%d.jpg", card.Id)))
	if err != nil {
		return err
	}

	var positionRect image.Rectangle
	if card.HasSubType("pendulum") {
		positionRect = BuildPositions.ArtworkPendulumBox
	} else {
		positionRect = BuildPositions.ArtworkBox
	}

	size := positionRect.Size()
	width, height := uint(size.X), uint(size.Y)

	artworkImage = resize.Resize(width, height, artworkImage, resize.Bilinear)
	imagesutils.DrawAt(img, artworkImage, positionRect.Min)

	return nil
}
