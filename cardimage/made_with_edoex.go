package cardimage

import (
	"edoex/environment"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutMadeWithEdoex(img draw.Image) error {
	edoexImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.TemplatesPath(), "made_with_edoex.png"))
	if err != nil {
		return err
	}

	imagesutils.DrawAt(img, edoexImage, BuildPositions.MadeWithEdoex)
	return nil
}
