package edopro

import (
	"edoex/environment"
	"edoex/models"
	"edoex/utils/filesutils"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Copies images from the images folder to the built pics folder
func CopyImages(cards []*models.Card) {
	for _, c := range cards {
		filename := fmt.Sprintf("%d.jpg", (*c).Id)
		path := filepath.Join(environment.ImagesPath(), filename)
		buildPath := filepath.Join(environment.BuildPath(), "pics", filename)

		content, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("Image '%s' does not exist", path)
			} else {
				log.Printf("Error reading '%s'", path)
			}

			continue
		}

		filesutils.WriteToFile(buildPath, content)
	}
}
