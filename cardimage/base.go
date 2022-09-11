package cardimage

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/imagesutils"
	"edoex/utils/sliceutils"
	"image"
	"path/filepath"
)

func GetCardBase(card *models.Card) (image.Image, error) {
	var baseFile string

	switch card.CardType {
	case "spell":
		baseFile = "spell.png"
	case "trap":
		baseFile = "trap.png"
	default:
		baseFile = getMonsterBaseFile(card.SubTypes)
	}

	basePath := filepath.Join(environment.TemplatesPath(), baseFile)
	image, err := imagesutils.LoadImageFromPath(basePath)
	if err != nil {
		return nil, err
	}

	return image, nil
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
