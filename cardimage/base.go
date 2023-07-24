package cardimage

import (
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/sliceutils"
	"image/draw"
	"path/filepath"
)

func GetCardBase(card *models.Card) (draw.Image, error) {
	logger.Verbosef("Getting base image for '%s' (%d)", card.Name, card.Id)
	var baseFile string

	switch card.CardType {
	case "spell":
		baseFile = "spell.png"
	case "trap":
		baseFile = "trap.png"
	default:
		baseFile = getMonsterBaseFile(card.SubTypes)
	}

	logger.Verbosef("%d - Base '%s'", card.Id, baseFile)
	basePath := filepath.Join(environment.GlobalTemplatesPath(), baseFile)
	image, err := imagesutils.LoadImageFromPath(basePath)
	if err != nil {
		return nil, err
	}

	return imagesutils.GetRGBA(image), nil
}

func getMonsterBaseFile(subTypes []string) string {
	specialTypes := []string{
		"xyz", "fusion", "ritual", "synchro", "token", "link",
	}

	for _, t := range subTypes {
		if sliceutils.Contains(specialTypes, t) {
			return t + ".png"
		}
	}

	if sliceutils.Contains(subTypes, "effect") {
		return "effect.png"
	} else {
		return "normal.png"
	}
}
