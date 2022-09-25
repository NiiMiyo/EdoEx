package cardimage

import (
	"edoex/environment"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutHologram(img draw.Image) (draw.Image, error) {
	hologramImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.TemplatesPath(), "hologram.png"))

	if err != nil {
		return nil, err
	}

	return imagesutils.DrawOver(img, hologramImage), nil
}
