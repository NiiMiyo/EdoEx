package parser

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"edoex/models"
	"edoex/utils/filesutils"
	"edoex/utils/sliceutils"

	"gopkg.in/yaml.v3"
)

type cardYaml struct {
	Id          int64    `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	CardType    string   `yaml:"card_type"`
	SubTypes    []string `yaml:"sub_types"`

	Atk            int64    `yaml:"atk"`
	Def            int64    `yaml:"def"`
	Level          int64    `yaml:"level"`
	Race           []string `yaml:"race"`
	Attribute      []string `yaml:"attribute"`
	PendulumEffect string   `yaml:"pendulum_effect"`
	LinkArrows     []string `yaml:"link_arrows"`
	Scale          int64    `yaml:"scale"`

	Ruleset  []string `yaml:"ruleset"`
	Alias    int64    `yaml:"alias"`
	Sets     []string `yaml:"sets"`
	Category []string `yaml:"category"`
	Strings  []string `yaml:"strings"`
}

var validCardTypes = []string{"monster", "spell", "trap"}

// todo: move 'skill' to main card type

// Parses and validates a YAML document to a Card struct.
// If it is not a valid Card returns `nil`
func CardFromYamlDocument(doc []byte, availableSets map[string]*models.Meta) (*models.Card, error) {
	var parsed cardYaml
	err := yaml.Unmarshal(doc, &parsed)
	if err != nil {
		return nil, errors.New("Parsing failed")
	}

	parsed.CardType = strings.ToLower(parsed.CardType)
	if !sliceutils.Contains(validCardTypes, parsed.CardType) {
		return nil, errors.New("Invalid type")
	}

	if parsed.Id <= 0 {
		return nil, errors.New("Invalid id")
	}

	if parsed.Name == "" {
		return nil, errors.New("Empty name")
	}

	if parsed.Description == "" {
		return nil, errors.New("Empty description")
	}

	for _, fieldPointer := range [](*[]string){
		&parsed.Race, &parsed.Ruleset, &parsed.Category, &parsed.SubTypes,
		&parsed.Attribute, &parsed.LinkArrows, &parsed.Sets,
	} {
		*fieldPointer = sliceutils.RemoveDuplicates(*fieldPointer)
		*fieldPointer = sliceutils.Map(*fieldPointer, strings.ToLower)
	}

	var cardSets []*models.Meta
	for _, s := range parsed.Sets {
		set, inExpansion := availableSets[s]

		if inExpansion {
			cardSets = append(cardSets, set)
			continue
		}

		code, err := strconv.ParseInt(s, 0, 64)
		if err == nil {
			set = getSetWithId(code, availableSets)

			if set != nil {
				cardSets = append(cardSets, set)
			} else {
				cardSets = append(cardSets, &models.Meta{
					Name:  "",
					Id:    code,
					Type:  models.MetaTypeSet,
					Alias: "",
				})
			}

			continue
		}

		log.Printf("Set '%s' on card '%s' (%d) does not exist\n", s, parsed.Name, parsed.Id)
	}

	return &models.Card{
		Id:                  parsed.Id,
		Name:                parsed.Name,
		Description:         parsed.Description,
		CardType:            parsed.CardType,
		SubTypes:            parsed.SubTypes,
		Atk:                 parsed.Atk,
		Def:                 parsed.Def,
		Level:               parsed.Level,
		Race:                parsed.Race,
		Attribute:           parsed.Attribute,
		PendulumDescription: parsed.PendulumEffect,
		LinkArrows:          parsed.LinkArrows,
		Scale:               parsed.Scale,
		Ruleset:             parsed.Ruleset,
		Alias:               parsed.Alias,
		Category:            parsed.Category,
		Strings:             parsed.Strings,
		Sets:                cardSets,
	}, nil
}

func CardsFromYamlFile(content []byte, availableSets map[string]*models.Meta) ([]*models.Card, error) {
	documents, err := filesutils.SplitYamlDocuments(content)
	if err != nil {
		return nil, err
	}

	var cards []*models.Card
	for _, d := range documents {
		parsed, err := CardFromYamlDocument(d, availableSets)
		if err != nil {
			return nil, err
		}

		cards = append(cards, parsed)
	}

	return cards, nil
}

func getSetWithId(id int64, availableSets map[string]*models.Meta) *models.Meta {
	for _, s := range availableSets {
		if id == s.Id {
			return s
		}
	}

	return nil
}
