package edopro

import (
	"os"
	"path/filepath"
	"strings"

	"edoex/environment"
	"edoex/logger"
	"edoex/utils/filesutils"
)

// Writes string.conf content
func BuildGlobalStrings() {
	var confStrings []string
	for _, m := range environment.MetasIds {
		conf := m.StringConfLine()

		if conf != "" {
			confStrings = append(confStrings, conf)
		}
	}

	fileContent := strings.Join(confStrings, "\n") + "\n"
	filesutils.WriteToFile(
		environment.BuildStringsPath(),
		[]byte(fileContent),
	)
}

func UpdateStrings() (string, error) {
	edoStringsPath := filepath.Join(environment.Config.Gamedir, "expansions/strings.conf")
	logger.Logf("Updating '%s'", edoStringsPath)

	oldStringsContent, err := os.ReadFile(edoStringsPath)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	newStringsContent, err := os.ReadFile(environment.BuildStringsPath())
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
