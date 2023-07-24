package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"image/draw"
	"path/filepath"
)

func PutMadeWithEdoex(img draw.Image, card *models.Card) error {
	logger.Verbosef("%d - Putting 'Made with EdoEx'", card.Id)
	edoexImage, err := imagesutils.LoadImageFromPath(
		filepath.Join(environment.GlobalTemplatesPath(), "made_with_edoex.png"))
	if err != nil {
		return err
	}

	imagesutils.DrawAt(img, edoexImage, BuildPositions.MadeWithEdoex)
	return nil
}
