package cmd

import (
	"edoex/edopro"
	"edoex/environment"
	"edoex/logger"
	"edoex/macro"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b", "compile"},
	Short:   "Builds the current expansion",
	Long:    `Builds the expansion source files in the current directory in the default way EDOPro will read them when importing a repository`,
	Run:     build,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func build(cmd *cobra.Command, args []string) {
	environment.UpdateConfig()

	logger.Logf("Building expansion '%s'", environment.Config.ExpansionName)

	logger.Logf("Preparing '%s' folder", environment.BuildDir)
	err := environment.ClearBuild()
	if err != nil {
		logger.ErrorErr("Error when cleaning previous build", err)
		return
	}

	err = environment.LoadExpansionData()
	if err != nil {
		return
	}

	logger.Log("Running macros")
	macro.ApplyMacros()

	logger.Logf("Writing '%s'", environment.BuildStringsPath())
	edopro.BuildGlobalStrings()

	logger.Logf("Writing '%s'", environment.BuildDatabasePath())
	err = edopro.WriteToCdb()
	if err != nil {
		logger.ErrorErr("Error when building database", err)
		return
	}

	logger.Log("Writing scripts")
	edopro.BuildScripts()

	logger.Log("Building images")
	edopro.BuildImages()
}
