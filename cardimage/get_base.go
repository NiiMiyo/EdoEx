package cardimage

import (
	"bufio"
	"edoex/environment"
	"edoex/models"
	"edoex/utils/sliceutils"
	"image"
	_ "image/png"
	"os"
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
	file, err := os.Open(basePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	image, _, err := image.Decode(reader)
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
