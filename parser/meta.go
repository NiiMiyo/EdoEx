package parser

import (
	"edoex/models"
	"edoex/utils/filesutils"
	"edoex/utils/sliceutils"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

type metaYaml struct {
	Id    int64  `yaml:"id"`
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Alias string `yaml:"alias"`
}

var validMetaTypes = []string{"counter", "set"}

// Parses and validates a YAML document to a Meta struct.
// If it is not a valid Meta returns `nil`
func MetaFromYamlDocument(doc []byte) (models.Meta, error) {
	var parsed metaYaml
	err := yaml.Unmarshal(doc, &parsed)
	if err != nil {
		return nil, errors.New("Parsing failed")
	}

	parsed.Type = strings.ToLower(parsed.Type)

	if !sliceutils.Contains(validMetaTypes, parsed.Type) {
		return nil, errors.New("Invalid type")
	}

	if parsed.Id <= 0 {
		return nil, errors.New("Invalid id")
	}

	if parsed.Name == "" {
		return nil, errors.New("Empty name")
	}

	switch parsed.Type {
	case "set":
		return &models.Set{
			Id:    parsed.Id,
			Name:  parsed.Name,
			Alias: parsed.Alias,
		}, nil

	case "counter":
		return &models.Counter{
			Id:   parsed.Id,
			Name: parsed.Name,
		}, nil
	}

	return nil, nil
}

// Parses a YAML file to an array of Meta structs.
// If at least one document on the file is not a valid Meta, returns `nil`
func MetaFromYamlFile(content []byte) ([]*models.Meta, error) {
	documents, err := filesutils.SplitYamlDocuments(content)
	if err != nil {
		return nil, err
	}

	var metas []*models.Meta
	for _, d := range documents {
		parsed, err := MetaFromYamlDocument(d)
		if err != nil {
			return nil, err
		}

		metas = append(metas, &parsed)
	}

	return metas, nil
}
