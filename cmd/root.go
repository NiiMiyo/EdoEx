package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"edoex/embedfiles"
	"edoex/environment/flags"
)

var rootCmd = &cobra.Command{
	Use:   "edoex",
	Short: "CLI tool to help creating expansions to EDOPro",
	Long:  embedfiles.EdoexLogo,
}

func init() {
	// todo: add verbose flag

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolVar(
		&flags.NoColor,
		"nocolor", false, "Output will not be colored",
	)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
