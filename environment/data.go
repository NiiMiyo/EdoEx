package environment

import (
	"log"
	"os"
	"strings"

	"edoex/models"
	"edoex/parser"
	"edoex/utils/filesutils"
)

// Filter to be used to get only YAML files
func isYamlFile(path string) bool {
	lower := strings.ToLower(path)
	return strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml")
}

// Walks through 'meta' folder and parses all yaml files
func loadMetaData() {
	log.Printf("Reading '%s' folder\n", SourceMetaDir)

	metaFiles, err := filesutils.WalkDirectoryAndFilter(SourceMetaPath(), isYamlFile)
	if err != nil {
		log.Fatalln(err)
	}

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
			current := MetasIds[m.Id]
			if current != nil {
				log.Printf(
					"Duplicated id '%d' - '%s' and '%s'. '%s' will be ignored",
					m.Id, m.Name, current.Name, m.Name)
				continue
			}

			MetasIds[m.Id] = m

			current = MetasAlias[m.AliasOrName()]
			if current == nil {
				MetasAlias[m.AliasOrName()] = m
			} else {
				log.Printf("Duplicated alias or name '%s'", m.AliasOrName())
			}
		}

		for _, m := range MetasIds {
			if m.Type == models.MetaTypeSet {
				Sets[m.AliasOrName()] = m
			}
		}
	}
}

// Loads all data from cards and metas
func LoadExpansionData() {
	loadMetaData()

	log.Printf("Reading '%s' folder\n", SourceCardsDir)

	cardFiles, err := filesutils.WalkDirectoryAndFilter(SourceCardsPath(), isYamlFile)
	if err != nil {
		log.Fatalln(err)
	}

	for _, path := range cardFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading '%s' - %s\n", path, err)
			continue
		}

		cards, err := parser.CardsFromYamlFile(content, Sets)
		if err != nil {
			log.Printf("Error parsing '%s' - %s\n", path, err)
			continue
		}

		for _, c := range cards {
			current := Cards[c.Id]
			if current == nil {
				Cards[c.Id] = c
			} else {
				log.Printf("Duplicated id '%d' - '%s' and '%s'. '%s' will be ignored",
					c.Id, c.Name, current.Name, c.Name)
			}
		}
	}
}
