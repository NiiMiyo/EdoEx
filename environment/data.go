package environment

import (
	"os"
	"strings"

	"edoex/logger"
	"edoex/models"
	"edoex/parser"
	"edoex/utils/filesutils"
)

// Filter to be used to get only YAML files
func isYamlFile(path string) bool {
	logger.Verbosef("Checking if file %s is valid YAML", path)
	lower := strings.ToLower(path)
	return strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml")
}

// Walks through 'meta' folder and parses all yaml files
func loadMetaData() error {
	logger.Logf("Reading '%s' folder", SourceMetaDir)

	logger.Verbosef("Walking '%s' folder", SourceMetaDir)
	metaFiles, err := filesutils.WalkDirectoryAndFilter(SourceMetaPath(), isYamlFile)
	if err != nil {
		logger.ErrorErr("Error loading metas", err)
		return err
	}

	for _, path := range metaFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			logger.ErrorfErr("Error parsing '%s'", err, path)
			continue
		}

		logger.Verbosef("Parsing metas from '%s'", path)
		metas, err := parser.MetaFromYamlFile(content)
		if err != nil {
			logger.ErrorfErr("Error parsing '%s'", err, path)
			continue
		}

		for _, m := range metas {
			current := MetasIds[m.Id]
			if current != nil {
				logger.Warnf(
					"Duplicated id '%d' for metas '%s' and '%s'. '%s' will be ignored",
					m.Id, m.Name, current.Name, m.Name)
				continue
			}

			MetasIds[m.Id] = m

			current = MetasAlias[m.AliasOrName()]
			if current == nil {
				MetasAlias[m.AliasOrName()] = m
			} else {
				logger.Warnf("Duplicated alias or name '%s'", m.AliasOrName())
			}
		}

		for _, m := range MetasIds {
			if m.Type == models.MetaTypeSet {
				Sets[m.AliasOrName()] = m
			}
		}
	}

	return nil
}

// Loads all data from cards and metas
func LoadExpansionData() error {
	err := loadMetaData()
	if err != nil {
		return err
	}

	logger.Logf("Reading '%s' folder", SourceCardsDir)

	logger.Verbosef("Walking '%s' folder", SourceCardsDir)
	cardFiles, err := filesutils.WalkDirectoryAndFilter(SourceCardsPath(), isYamlFile)
	if err != nil {
		logger.ErrorErr("Error loading cards", err)
		return err
	}

	for _, path := range cardFiles {
		logger.Verbosef("Reading card file '%s'", path)
		content, err := os.ReadFile(path)
		if err != nil {
			logger.ErrorfErr("Error reading '%s'", err, path)
			continue
		}

		logger.Verbosef("Parsing card file '%s'", path)
		cards, err := parser.CardsFromYamlFile(content, Sets)
		if err != nil {
			logger.ErrorfErr("Error parsing '%s'", err, path)
			continue
		}

		for _, c := range cards {
			current := Cards[c.Id]
			if current == nil {
				Cards[c.Id] = c
			} else {
				logger.Warnf("Duplicated id '%d' for cards '%s' and '%s'. '%s' will be ignored",
					c.Id, c.Name, current.Name, c.Name)
			}
		}
	}

	return nil
}
