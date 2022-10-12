package cmd

import (
	"path/filepath"

	"edoex/edopro"
	"edoex/environment"
	"edoex/logger"
	"edoex/utils/filesutils"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s", "synchro", "synchronize"},
	Short:   "Syncs the current expansion with your game folder",
	Long: `Builds the current expansions and syncs it to the EDOPro installation specified on edoex.config.yaml "gamedir" property.
If not specified on the current project, will use the one on the EdoEx installation folder.
If both are not specified the command fails.`,
	Run: sync,
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func sync(cmd *cobra.Command, args []string) {
	buildCmd.Run(cmd, args)

	logger.Log()

	if environment.Config.Gamedir == "" {
		logger.Errorf("Property 'gamedir' not defined on '%s'", environment.ConfigFile)
		return
	}

	logger.Logf(
		"Preparing to sync '%s' with EDOPro in '%s'",
		environment.Config.ExpansionName,
		environment.Config.Gamedir,
	)

	logger.Logf("Updating '%s'", environment.Config.ExpansionSyncPath())
	err := filesutils.ZipFiles(environment.Config.ExpansionSyncPath(), filesToZip())
	if err != nil {
		logger.ErrorErr("Error syncing files", err)
		return
	}

	newStrings, err := edopro.UpdateStrings()
	if err != nil {
		logger.ErrorErr("Error generating strings.conf", err)
		return
	}

	edoStringsPath := filepath.Join(environment.Config.Gamedir, "expansions/strings.conf")
	filesutils.WriteToFile(edoStringsPath, []byte(newStrings))
}

// Returns a map where the key is which file/folder should be zipped and the
// value is the path in the .zip file
func filesToZip() map[string]string {
	files := make(map[string]string)

	files[environment.BuildDatabasePath()] = environment.Config.ExpansionName + ".cdb"

	scripts := filepath.Join(environment.BuildPath(), "script")
	files[scripts] = "script"

	pics := filepath.Join(environment.BuildPath(), "pics")
	files[pics] = "pics"

	return files
}
