/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"edoex/embedfiles"
)

var rootCmd = &cobra.Command{
	Use:   "edoex",
	Short: "CLI tool to help creating expansions to EDOPro",
	Long:  embedfiles.EdoexLogo,
}

// todo: add verbose flag
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
