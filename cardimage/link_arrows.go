package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image"
	"path/filepath"
)

var linkArrowsPath = filepath.Join(environment.TemplatesPath(), "link_arrows")

func PutLinkArrows(img image.Image, card *models.Card) (image.Image, error) {
	if !card.HasSubType("link") {
		return img, nil
	}

	disabledArrowsImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(linkArrowsPath, "disabled.png"))
	if err != nil {
		return nil, err
	}
	img = imagesutils.DrawOver(img, disabledArrowsImage)

	for _, direction := range card.LinkArrows {
		directionImage, err := imagesutils.LoadImageFromPath(
			filepath.Join(linkArrowsPath, direction+".png"))
		if err != nil {
			return nil, err
		}

		img = imagesutils.DrawOver(img, directionImage)
	}

	return img, nil
}
