package environment

import (
	"edoex/models"
	"edoex/parser"
	"edoex/utils/filesutils"
	"fmt"
	"log"
	"os"
	"strings"
)

// Filter to be used to get only YAML files
func isYamlFile(path string) bool {
	lower := strings.ToLower(path)
	return strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml")
}

// Walks through 'meta' folder and parses all yaml files
func getMetaData() []models.Meta {
	fmt.Println()
	log.Printf("Reading '%s' folder\n", MetaDir)

	metaFiles, err := filesutils.WalkDirectoryAndFilter(MetaPath(), isYamlFile)
	if err != nil {
		log.Fatalln(err)
	}

	var allMetas []models.Meta
	for _, path := range metaFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error parsing '%s' - %s\n", path, err)
			continue
		}

		metas, err := parser.MetaFromYamlFile(content)
		if err != nil {
			log.Printf("Error parsing '%s' - %s\n", path, err)
			continue
		}

		for _, m := range metas {
			allMetas = append(allMetas, m)
		}
	}

	return allMetas
}

// Gets all data from cards and metas
func GetExpansionData() (cards []models.Card, meta []models.Meta) {
	metas := getMetaData()

	// todo: get cards

	return nil, metas
}
