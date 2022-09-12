package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"fmt"
	"image"
	"image/draw"
	"path/filepath"

	"github.com/nfnt/resize"
)

var (
	normalCardPosition   = image.Rect(107, 210, 633, 736)
	pendulumCardPosition = image.Rect(68, 206, 670, 654)
)

func PutArtwork(img image.Image, card *models.Card) (image.Image, error) {
	artworkImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.ArtworksPath(), fmt.Sprintf("%d.jpg", card.Id)))
	if err != nil {
		return nil, err
	}

	var positionRect image.Rectangle
	if card.HasSubType("pendulum") {
		positionRect = pendulumCardPosition
	} else {
		positionRect = normalCardPosition
	}

	size := positionRect.Size()
	width, height := uint(size.X), uint(size.Y)

	artworkImage = resize.Resize(width, height, artworkImage, resize.Bilinear)
	drawableArtwork := imagesutils.GetRGBA(img)

	draw.Draw(drawableArtwork, positionRect, artworkImage, image.Point{X: 0, Y: 0}, draw.Over)
	return drawableArtwork, nil
}
