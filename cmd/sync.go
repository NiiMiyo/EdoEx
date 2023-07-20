package cmd

import (
	"fmt"
	"path/filepath"

	"edoex/edopro"
	"edoex/environment"
	"edoex/environment/flags"
	"edoex/logger"
	"edoex/utils/filesutils"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s", "synchro", "synchronize"},
	Short:   "Syncs the current expansion with your game folder",
	Long: `Builds the current expansions and syncs it to the simulator installation specified on edoex.config.yaml "edopro_path" or "omega_path" property.
If not specified on the current project, will use the one on the EdoEx installation folder.
If both are not specified the command fails.`,
	Run: sync,
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.PersistentFlags().StringVar(
		&flags.Simulator,
		"simulator", "", "Simulator used to build. Either \"edopro\" or \"omega\"",
	)
}

func sync(cmd *cobra.Command, args []string) {
	buildCmd.Run(cmd, args)

	logger.Log()

	simulator, err := environment.GetSimulator()
	if err != nil {
		logger.Error(err)
		return
	}

	switch simulator {
	case environment.EdoproSimulator:
		SyncEdopro()
	case environment.OmegaSimulator:
		SyncOmega()
	}
}

func SyncEdopro() error {
	if environment.Config.EdoproPath == "" {
		return fmt.Errorf("Property 'edopro_path' not defined on '%s'", environment.ConfigFile)
	}

	logger.Logf(
		"Preparing to sync '%s' with EDOPro in '%s'",
		environment.Config.ExpansionName,
		environment.Config.EdoproPath,
	)

	logger.Logf("Updating '%s'", environment.Config.EdoproSyncZipPath())
	err := filesutils.ZipFiles(environment.Config.EdoproSyncZipPath(), edoproFilesToZip())
	if err != nil {
		return fmt.Errorf("Error syncing files", err)
	}

	newStrings, err := edopro.UpdateStrings()
	if err != nil {
		return fmt.Errorf("Error generating strings.conf", err)
	}

	edoStringsPath := filepath.Join(environment.Config.EdoproPath, "expansions/strings.conf")
	filesutils.WriteToFile(edoStringsPath, []byte(newStrings))

	return nil
}

func SyncOmega() error {
	if environment.Config.OmegaPath == "" {
		return fmt.Errorf("Property 'omega_path' not defined on '%s'", environment.ConfigFile)
	}

	logger.Logf(
		"Preparing to sync '%s' with YGO Omega in '%s'",
		environment.Config.ExpansionName,
		environment.Config.OmegaPath,
	)

	logger.Log("Syncing images")
	omegaImagesPath := filepath.Join(environment.Config.OmegaSyncPath(), "Arts")
	err := filesutils.CopyDirectoryContent(environment.OmegaImagesBuildPath(), omegaImagesPath)
	if err != nil {
		return err
	}

	logger.Log("Syncing scripts")
	omegaScriptsPath := filepath.Join(environment.Config.OmegaSyncPath(), "Scripts")
	err = filesutils.CopyDirectoryContent(environment.OmegaScriptsBuildPath(), omegaScriptsPath)
	if err != nil {
		return err
	}

	logger.Log("Syncing database")
	omegaDatabasePath := filepath.Join(environment.Config.OmegaSyncPath(), "Databases", environment.Config.ExpansionName+".db")
	err = filesutils.CopyFile(environment.OmegaDatabaseBuildPath(), omegaDatabasePath)
	if err != nil {
		return err
	}

	return nil
}

// Returns a map where the key is which file/folder should be zipped and the
// value is the path in the .zip file
func edoproFilesToZip() map[string]string {
	files := make(map[string]string)

	files[environment.EdoproDatabaseBuildPath()] = environment.Config.ExpansionName + ".cdb"

	scripts := filepath.Join(environment.BuildPath(), "script")
	files[scripts] = "script"

	pics := filepath.Join(environment.BuildPath(), "pics")
	files[pics] = "pics"

	return files
}
