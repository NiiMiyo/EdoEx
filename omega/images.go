package omega

import (
	"fmt"
	"os"
	"path/filepath"

	"edoex/environment"
	"edoex/logger"
	"edoex/utils/filesutils"
)

func BuildImages() {
	os.MkdirAll(environment.OmegaImagesBuildPath(), os.ModeDir)
	for _, card := range environment.Cards {
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

		imageFilename := fmt.Sprintf("%d.jpg", card.Id)
		cardImagePath := filepath.Join(environment.OmegaImagesBuildPath(), imageFilename)

		filesutils.CopyFile(artworkPath, cardImagePath)
	}
}
