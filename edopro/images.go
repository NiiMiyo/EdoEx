package edopro

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
