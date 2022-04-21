package edopro

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"edoex/environment"
	"edoex/models"
)

// Returns strings.conf content
func BuildGlobalStrings(metas []*models.Meta) string {
	var confStrings []string
	for _, m := range metas {
		confStrings = append(confStrings, (*m).StringConfLine())
	}

	return strings.Join(confStrings, "\n") + "\n"
}

func UpdateStrings() (string, error) {
	edoStringsPath := filepath.Join(environment.Config.Gamedir, "expansions/strings.conf")
	log.Printf("Updating '%s'\n", edoStringsPath)

	oldStringsContent, err := os.ReadFile(edoStringsPath)
	if err != nil {
		return "", err
	}

	newStringsContent, err := os.ReadFile(environment.StringsPath())
	if err != nil {
		return "", err
	}

	var stringsConf []string
	newStrings := strings.Split(string(newStringsContent), "\n")
	oldStrings := strings.Split(string(oldStringsContent), "\n")

	lineMatches := make(map[int]interface{})

	// update codes that are on both strings.conf
	for _, oldLine := range oldStrings {
		oldLine = strings.Trim(oldLine, " ")
		if oldLine == "" {
			continue
		}

		matches := false

		// gets "!set 0xcode" or "!counter 0xcode"
		check := strings.Join(strings.Split(oldLine, " ")[:2], " ")

		// Checks if there is a new line with the same code from an old line
		for i, newLine := range newStrings {
			if strings.HasPrefix(newLine, check) {
				stringsConf = append(stringsConf, newLine)
				lineMatches[i] = struct{}{}
				matches = true
				break
			}
		}

		if !matches {
			stringsConf = append(stringsConf, oldLine)
		}
	}

	// Add codes that aren't on EDOPro's strings.conf
	for i, newLine := range newStrings {
		_, contains := lineMatches[i]
		if !contains {
			stringsConf = append(stringsConf, newLine)
		}
	}

	return strings.Join(stringsConf, "\n"), nil
}
