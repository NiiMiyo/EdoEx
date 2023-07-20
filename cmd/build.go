package cmd

import (
	"edoex/edopro"
	"edoex/environment"
	"edoex/environment/flags"
	"edoex/logger"
	"edoex/macro"
	"edoex/omega"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b", "compile"},
	Short:   "Builds the current expansion",
	Long:    `Builds the expansion source files in the current directory in the default way your simulator will read them when importing a repository`,
	Run:     build,
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().StringVar(
		&flags.Simulator,
		"simulator", "", "Simulator used to build. Either \"edopro\" or \"omega\"",
	)
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

	simulator, err := environment.GetSimulator()
	if err != nil {
		logger.Error(err)
		return
	}

	switch simulator {
	case environment.EdoproSimulator:
		BuildEdopro()

	case environment.OmegaSimulator:
		BuildOmega()
	}
}

func BuildOmega() {
	logger.Logf("Writing '%s'", environment.OmegaDatabaseBuildPath())
	err := omega.WriteToDb()
	if err != nil {
		logger.ErrorErr("Error when building database", err)
		return
	}

	logger.Log("Building scripts")
	omega.BuildScripts()

	logger.Log("Building images")
	omega.BuildImages()
}

func BuildEdopro() {
	logger.Logf("Writing '%s'", environment.EdoproStringsBuildPath())
	edopro.BuildGlobalStrings()

	logger.Logf("Writing '%s'", environment.EdoproDatabaseBuildPath())
	err := edopro.WriteToCdb()
	if err != nil {
		logger.ErrorErr("Error when building database", err)
		return
	}

	logger.Log("Writing scripts")
	edopro.BuildScripts()

	logger.Log("Building images")
	edopro.BuildImages()
}
