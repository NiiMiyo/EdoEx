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
func getMetaData() []*models.Meta {
	log.Printf("Reading '%s' folder\n", MetaDir)

	metaFiles, err := filesutils.WalkDirectoryAndFilter(MetaPath(), isYamlFile)
	if err != nil {
		log.Fatalln(err)
	}

	var allMetas []*models.Meta
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

// Walks through 'cards' folder and parses all yaml files
func getCardsData(metas []*models.Meta) []*models.Card {
	log.Printf("Reading '%s' folder\n", CardsDir)

	cardFiles, err := filesutils.WalkDirectoryAndFilter(CardsPath(), isYamlFile)
	if err != nil {
		log.Fatalln(err)
	}

	sets := make(map[string]*models.Set)
	for _, m := range metas {
		s, ok := (*m).(*models.Set)

		if ok {
			if s.Alias != "" {
				sets[s.Alias] = s
			} else {
				sets[s.Name] = s
			}
		}
	}

	var allCards []*models.Card
	for _, path := range cardFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error parsing '%s' = %s\n", path, err)
			continue
		}

		cards, err := parser.CardsFromYamlFile(content, sets)
		if err != nil {
			log.Printf("Error parsing '%s' - %s\n", path, err)
			continue
		}

		for _, c := range cards {
			allCards = append(allCards, c)
		}

	}
	return allCards
}

// Gets all data from cards and metas
func GetExpansionData() ([]*models.Card, []*models.Meta) {
	metas := getMetaData()

	return getCardsData(metas), metas
}
