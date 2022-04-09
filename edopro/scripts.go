package edopro

import (
	"edoex/embedfiles"
	"edoex/environment"
	"edoex/models"
	"edoex/utils/filesutils"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Default content for cID.lua when the card has no script on scripts folder
func DefaultScript(card *models.Card) []byte {
	return []byte(fmt.Sprintf(embedfiles.DefaultCardScript, (*card).Id, (*card).Name))
}

// Copies/creates cID.lua files to scripts folder
func BuildScripts(cards []*models.Card) {
	for _, c := range cards {
		filename := fmt.Sprintf("c%d.lua", (*c).Id)
		path := filepath.Join(environment.ScriptsPath(), filename)
		buildPath := filepath.Join(environment.BuildPath(), "script", filename)

		content, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("Script '%s' does not exist.\n", path)
			} else {
				log.Printf("Error reading '%s'", path)
			}

			content = DefaultScript(c)
		}

		filesutils.WriteToFile(buildPath, content)
	}
}
