package omega

import (
	"fmt"
	"os"
	"path/filepath"

	"edoex/embedfiles"
	"edoex/environment"
	"edoex/logger"
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
		path := filepath.Join(environment.SourceScriptsPath(), filename)
		buildPath := filepath.Join(environment.OmegaScriptsBuildPath(), filename)

		err := filesutils.CopyFile(path, buildPath)
		if err != nil {
			if os.IsNotExist(err) {
				logger.Warnf("Script '%s' does not exist", path)
			} else {
				logger.ErrorfErr("Error reading '%s'", err, path)
			}

			filesutils.WriteToFile(buildPath, DefaultScript(c))
		}
	}
}
