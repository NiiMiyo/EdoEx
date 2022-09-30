package imagesutils

import (
	"image"
	"image/color"

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
