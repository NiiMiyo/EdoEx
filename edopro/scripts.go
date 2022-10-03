package edopro

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"edoex/embedfiles"
	"edoex/environment"
	"edoex/models"
	"edoex/utils/filesutils"
)

// Default content for cID.lua when the card has no script on scripts folder
func DefaultScript(card *models.Card) []byte {
	return []byte(fmt.Sprintf(embedfiles.DefaultCardScript, (*card).Id, (*card).Name))
}

// Copies/creates cID.lua files to scripts folder
func BuildScripts() {
	for _, c := range environment.Cards {
		filename := fmt.Sprintf("c%d.lua", (*c).Id)
		path := filepath.Join(environment.ScriptsPath(), filename)
		buildPath := filepath.Join(environment.BuildPath(), "script", filename)

		err := filesutils.CopyFile(path, buildPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("Script '%s' does not exist.\n", path)
			} else {
				log.Printf("Error reading '%s'", path)
			}

			filesutils.WriteToFile(buildPath, DefaultScript(c))
		}
	}
}
