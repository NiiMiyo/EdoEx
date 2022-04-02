package cmd

import (
	"fmt"
	"log"

	"edoex/embedfiles"
	"edoex/environment"
	"edoex/utils/filesutils"
	"edoex/utils/foldersutils"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init <expansion>",
	Aliases: []string{"initialize", "i", "new"},
	Short:   "Creates an expansion with base files",
	Long:    "Creates a new folder in current directory with expansion name and base files",
	Run:     initialize,
	Args:    cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initialize(cmd *cobra.Command, args []string) {
	// todo: template flag to generate cards and meta

	expansionName := args[0]

	empty, err := foldersutils.IsEmpty(environment.WorkingDir)
	if err != nil {
		log.Fatalln(err)
	}
	if !empty {
		log.Fatalf("Current directory '%s' is not empty\n", environment.WorkingDir)
	}

	log.Printf("Initializing expansion '%s'\n", expansionName)

	files := defaultFiles(expansionName)
	for _, f := range files {
		log.Printf("Creating %s file\n", f.Name)

		err = filesutils.WriteToFile(f.Path, f.Content)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

type initFile struct {
	Name    string
	Path    string
	Content []byte
	// todo: isTemplate bool
}

func defaultFiles(expansionName string) (files []initFile) {
	// todo: isTemplate param

	configContent := fmt.Sprintf(embedfiles.DefaultExpansionConfig, expansionName)
	configFile := initFile{
		Name:    "expansion configuration",
		Path:    "./edoex.config.yaml",
		Content: []byte(configContent),
	}

	files = append(files, configFile)
	return files
}
