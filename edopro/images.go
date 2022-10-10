package edopro

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"edoex/cardimage"
	"edoex/environment"
	"edoex/utils/filesutils"
)

func BuildImages() {
	os.MkdirAll(environment.BuildPicsPath(), os.ModeDir)

	for _, c := range environment.Cards {
		artworkPath := filepath.Join(environment.SourceArtworksPath(), fmt.Sprintf("%d.jpg", c.Id))
		hasArtwork, err := filesutils.Exists(artworkPath)

		if err != nil {
			log.Printf("Error accessing artwork of '%d':%v\n", c.Id, err)
			continue
		}

		if !hasArtwork {
			log.Printf("WARN '%d' has no artwork.\n", c.Id)
			continue
		}

		cardImage, err := cardimage.BuildCardImage(c)
		if err != nil {
			log.Printf("Error building '%d':%v\n", c.Id, err)
			continue
		}

		imageFilename := fmt.Sprintf("%d.jpg", c.Id)
		cardImagePath := filepath.Join(environment.BuildPicsPath(), imageFilename)

		file, err := os.Create(cardImagePath)
		if err != nil {
			log.Printf("Error saving '%d.jpg'\n", c.Id)
			continue
		}
		defer file.Close()

		jpeg.Encode(file, cardImage, nil)
	}
}
