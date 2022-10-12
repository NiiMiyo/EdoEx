package edopro

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"edoex/cardimage"
	"edoex/environment"
	"edoex/logger"
	"edoex/utils/filesutils"
)

func BuildImages() {
	os.MkdirAll(environment.BuildPicsPath(), os.ModeDir)

	for _, c := range environment.Cards {
		artworkPath := filepath.Join(environment.SourceArtworksPath(), fmt.Sprintf("%d.jpg", c.Id))
		hasArtwork, err := filesutils.Exists(artworkPath)

		if err != nil {
			logger.ErrorfErr("Error accessing artwork of '%d'", err, c.Id)
			continue
		}

		if !hasArtwork {
			logger.Warnf("Card '%d' has no artwork", c.Id)
			continue
		}

		cardImage, err := cardimage.BuildCardImage(c)
		if err != nil {
			logger.ErrorfErr("Error building '%d'", err, c.Id)
			continue
		}

		imageFilename := fmt.Sprintf("%d.jpg", c.Id)
		cardImagePath := filepath.Join(environment.BuildPicsPath(), imageFilename)

		file, err := os.Create(cardImagePath)
		if err != nil {
			logger.Errorf("Error saving '%d.jpg'", c.Id)
			continue
		}
		defer file.Close()

		jpeg.Encode(file, cardImage, nil)
	}
}
