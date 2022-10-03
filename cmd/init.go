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
	Long:    "Creates a new expansion in current directory with expansion name and base files",
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

// Returns which files should be created on init command
func defaultFiles(expansionName string) (files []initFile) {
	// todo: isTemplate param

	configContent := fmt.Sprintf(embedfiles.DefaultExpansionConfig, expansionName)
	configFile := initFile{
		Name:    "expansion configuration",
		Path:    environment.ConfigFile,
		Content: []byte(configContent),
	}

	cardsReadmeFile := initFile{
		Name:    "cards readme",
		Path:    "./cards/readme.md.txt",
		Content: []byte(embedfiles.CardsReadme),
	}

	metaReadmeFile := initFile{
		Name:    "meta readme",
		Path:    "./meta/readme.md.txt",
		Content: []byte(embedfiles.MetaReadme),
	}

	scriptsReadmeFile := initFile{
		Name:    "scripts readme",
		Path:    "./scripts/readme.md.txt",
		Content: []byte(embedfiles.ScriptsReadme),
	}

	artworksReadmeFile := initFile{
		Name:    "artworks readme",
		Path:    "./artworks/readme.md.txt",
		Content: []byte(embedfiles.ArtworksReadme),
	}

	macrosReadmeFile := initFile{
		Name:    "macros readme",
		Path:    "./cards/macros.md.txt",
		Content: []byte(embedfiles.MacrosReadme),
	}

	files = append(files, configFile)
	files = append(files, cardsReadmeFile)
	files = append(files, metaReadmeFile)
	files = append(files, scriptsReadmeFile)
	files = append(files, artworksReadmeFile)
	files = append(files, macrosReadmeFile)
	return files
}
