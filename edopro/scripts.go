package edopro

import (
	"fmt"
	"os"
	"path/filepath"

	"edoex/embedfiles"
	"edoex/environment"
	"edoex/logger"
	"edoex/models"
	"edoex/utils/filesutils"
	"edoex/utils/luautils"
)

// Default content for cID.lua when the card has no script on scripts folder
func DefaultScript(card *models.Card) []byte {
	return []byte(fmt.Sprintf(embedfiles.DefaultCardScript, (*card).Id, (*card).Name))
}

// Copies/creates cID.lua files to scripts folder
func BuildScripts() {
	for _, c := range environment.Cards {
		logger.Verbosef("Building script for '%s' (%d)", c.Name, c.Id)
		filename := fmt.Sprintf("c%d.lua", (*c).Id)
		path := filepath.Join(environment.SourceScriptsPath(), filename)
		buildPath := filepath.Join(environment.EdoproScriptsBuildPath(), filename)

		useDefault := false
		exists, err := filesutils.Exists(path)
		if err != nil {
			logger.ErrorfErr("Error reading '%s'", err, path)
			useDefault = true

		} else if !exists {
			logger.Warnf("Script '%s' does not exist", path)
			useDefault = true

		} else {
			content, err := os.ReadFile(path)
			if err != nil {
				useDefault = true
			} else {
				logger.Verbosef("Minifying '%s'", path)
				filesutils.WriteToFile(buildPath, luautils.MinifyCode(content))
			}
		}

		if useDefault {
			logger.Verbose("Writing default code")
			filesutils.WriteToFile(buildPath, DefaultScript(c))
		}
	}
}
