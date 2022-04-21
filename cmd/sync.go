package cmd

import (
	"edoex/edopro"
	"edoex/environment"
	"edoex/utils/filesutils"
	"fmt"
	"log"
	"path/filepath"

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

	fmt.Println()

	if environment.Config.Gamedir == "" {
		log.Fatalf(
			"Property 'gamedir' not defined on '%s'\n",
			environment.ConfigFile,
		)
	}

	log.Printf(
		"Preparing to sync '%s' with EDOPro in '%s'\n",
		environment.Config.ExpansionName,
		environment.Config.Gamedir,
	)

	log.Printf("Updating '%s'\n", environment.Config.ExpansionSyncPath())
	err := filesutils.ZipFiles(environment.Config.ExpansionSyncPath(), filesToZip())
	if err != nil {
		log.Fatalln(err)
	}

	newStrings, err := edopro.UpdateStrings()
	if err != nil {
		log.Fatalln(err)
	}

	edoStringsPath := filepath.Join(environment.Config.Gamedir, "expansions/strings.conf")
	filesutils.WriteToFile(edoStringsPath, []byte(newStrings))
}

// Returns a map where the key is which file/folder should be zipped and the
// value is the path in the .zip file
func filesToZip() map[string]string {
	files := make(map[string]string)

	files[environment.DatabasePath()] = environment.Config.ExpansionName + ".cdb"

	scripts := filepath.Join(environment.BuildPath(), "script")
	files[scripts] = "script"

	pics := filepath.Join(environment.BuildPath(), "pics")
	files[pics] = "pics"

	return files
}
