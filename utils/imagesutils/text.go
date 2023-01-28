package imagesutils

import (
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func GetFontFace(fontBytes []byte, size float64) (font.Face, error) {
	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}

	return face, nil
}

func TransparentBackgroundText(
	text string, textColor color.Color, fontFace font.Face) image.Image {

	cwh := gg.NewContext(0, 0)
	cwh.SetFontFace(fontFace)
	w, h := cwh.MeasureString(text)

	c := gg.NewContext(int(w), int(h))
	c.SetRGBA(0, 0, 0, 0)
	c.Clear()

	c.SetColor(textColor)
	c.SetFontFace(fontFace)
	c.DrawStringAnchored(text, 0, 0, 0, 0.75)
	return c.Image()
}

func JustifiedText(text string, textColor color.Color, fontFace font.Face,
	width int) image.Image {

	cwh := gg.NewContext(0, 0)
	cwh.SetFontFace(fontFace)
	_, lnh := cwh.MeasureString(text[:1])
	lineHeight := int(math.Ceil(lnh))

	lines := wordWrap(text, cwh, width)
	textBox := image.NewRGBA(image.Rect(0, 0, width, lineHeight*len(lines)))

	for i, ln := range lines {
		var lineImage image.Image

		if !ln.wrapped {
			lineImage = TransparentBackgroundText(ln.text, textColor, fontFace)
		} else {
			drawable := image.NewRGBA(image.Rect(0, 0, width, lineHeight))
			words, spaceInBetween := justifyLine(ln.text, cwh, width)
			offset := 0

			for _, w := range words {
				wordImg := TransparentBackgroundText(w, textColor, fontFace)
				DrawAt(drawable, wordImg, image.Point{offset, 0})
				offset += wordImg.Bounds().Dx() + spaceInBetween
			}

			lineImage = drawable
		}

		DrawAt(textBox, lineImage, image.Point{0, i * lineHeight})
	}

	return textBox
}
