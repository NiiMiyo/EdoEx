package edopro

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"sync"

	"edoex/cardimage"
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/filesutils"
)

func BuildImages() {
	os.MkdirAll(environment.EdoproImagesBuildPath(), os.ModeDir)
	var wg sync.WaitGroup

	for _, c := range environment.Cards {
		wg.Add(1)
		go func(card *models.Card) {
			defer wg.Done()

			artworkPath := filepath.Join(environment.SourceArtworksPath(), fmt.Sprintf("%d.jpg", card.Id))
			hasArtwork, err := filesutils.Exists(artworkPath)

			if err != nil {
				logger.ErrorfErr("Error accessing artwork of '%d'", err, card.Id)
				return
			}

			if !hasArtwork {
				logger.Warnf("Card '%d' has no artwork", card.Id)
				return
			}

			cardImage, err := cardimage.BuildCardImage(card)
			if err != nil {
				logger.ErrorfErr("Error building '%d'", err, card.Id)
				return
			}

			imageFilename := fmt.Sprintf("%d.jpg", card.Id)
			cardImagePath := filepath.Join(environment.EdoproImagesBuildPath(), imageFilename)

			file, err := os.Create(cardImagePath)
			if err != nil {
				logger.Errorf("Error saving '%d.jpg'", card.Id)
				return
			}
			defer file.Close()

			jpeg.Encode(file, cardImage, nil)
		}(c)
	}

	wg.Wait()
}
