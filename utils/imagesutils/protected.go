package imagesutils

import (
	"math"
	"strings"

	"github.com/fogleman/gg"
)

func wordWrap(text string, cwh *gg.Context, maxWidth int) []wordWrapResponse {
	lines := make([]wordWrapResponse, 0)

	for _, fullLine := range strings.Split(text, "\n") {
		safeLine := ""
		for _, word := range strings.Split(fullLine, " ") {
			if safeLine == "" {
				safeLine += word
				continue
			}

			checkLine := safeLine + " " + word
			checkLineWidth, _ := cwh.MeasureString(checkLine)

			if int(math.Ceil(checkLineWidth)) > maxWidth {
				lines = append(lines, wordWrapResponse{
					text:    safeLine,
					wrapped: true,
				})

				safeLine = word
			} else {
				safeLine = checkLine
				continue
			}
		}

		if safeLine != "" {
			lines = append(lines, wordWrapResponse{
				text:    safeLine,
				wrapped: false,
			})
		}
	}

	return lines
}

func justifyLine(text string, cwh *gg.Context, maxWidth int) ([]string, int) {
	words := strings.Split(text, " ")
	joinnedWords := strings.Replace(text, " ", "", -1)
	width, _ := cwh.MeasureString(joinnedWords)

	return words, int(math.Floor((float64(maxWidth) - width) / float64(len(words)-1)))
}

type wordWrapResponse struct {
	text    string
	wrapped bool
}
