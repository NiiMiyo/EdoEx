package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

var linkArrowsPath = filepath.Join(environment.TemplatesPath(), "link_arrows")

func PutLinkArrows(img draw.Image, card *models.Card) error {
	if !card.HasSubType("link") {
		return nil
	}

	disabledArrowsImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(linkArrowsPath, "disabled.png"))
	if err != nil {
		return err
	}
	imagesutils.DrawAt(img, disabledArrowsImage, BuildPositions.linkArrows)

	for _, direction := range card.LinkArrows {
		directionImage, err := imagesutils.LoadImageFromPath(
			filepath.Join(linkArrowsPath, direction+".png"))
		if err != nil {
			return err
		}

		imagesutils.DrawAt(img, directionImage, BuildPositions.linkArrows)
	}

	return nil
}
