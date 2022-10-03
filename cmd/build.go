package cmd

import (
	"log"
	"os"

	"edoex/edopro"
	"edoex/environment"
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

	log.Printf("Building expansion '%s'\n", environment.Config.ExpansionName)

	log.Printf("Preparing '%s' folder\n", environment.BuildDir)
	os.RemoveAll(environment.BuildPath())

	environment.LoadExpansionData()

	log.Println("Running macros")
	macro.ApplyMacros()

	log.Printf("Writing '%s'", environment.StringsPath())
	edopro.BuildGlobalStrings()

	log.Printf("Writing '%s'", environment.DatabasePath())
	err := edopro.WriteToCdb()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Writing scripts")
	edopro.BuildScripts()

	log.Println("Building images")
	edopro.BuildImages()
}
