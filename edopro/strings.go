package edopro

import (
	"edoex/models"
	"strings"
)

// Returns strings.conf content
func BuildGlobalStrings(metas ...models.Meta) string {
	var confStrings []string
	for _, m := range metas {
		confStrings = append(confStrings, m.StringConfLine())
	}

	return strings.Join(confStrings, "\n") + "\n"
}
