package edopro

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"edoex/cardimage"
	"edoex/environment"
	"edoex/models"
	"edoex/utils/filesutils"
)

// Copies images from the images folder to the built pics folder
func CopyImages(cards []*models.Card) {
	for _, c := range cards {
		filename := fmt.Sprintf("%d.jpg", (*c).Id)
		path := filepath.Join(environment.ImagesPath(), filename)
		buildPath := filepath.Join(environment.BuildPath(), "pics", filename)

		err := filesutils.CopyFile(path, buildPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("Image '%s' does not exist", path)
			} else {
				log.Printf("Error reading '%s'", path)
			}
		}
	}
}

func BuildImages(cards []*models.Card) {
	os.MkdirAll(environment.PicsPath(), os.ModeDir)

	for _, c := range cards {
		artworkPath := filepath.Join(environment.ArtworksPath(), fmt.Sprintf("%d.jpg", c.Id))
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
		cardImagePath := filepath.Join(environment.PicsPath(), imageFilename)

		file, err := os.Create(cardImagePath)
		if err != nil {
			log.Printf("Error saving '%d.jpg'\n", c.Id)
			continue
		}
		defer file.Close()

		jpeg.Encode(file, cardImage, nil)
	}
}
