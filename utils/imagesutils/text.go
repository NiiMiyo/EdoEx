package imagesutils

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func GetFontFace(fontBytes []byte, size, dpi float64) (font.Face, error) {
	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}

	return face, nil
}

func TransparentBackgroundText(
	text string, textColor color.Color, fontFace font.Face, w, h int) image.Image {

	c := gg.NewContext(w, h)
	c.SetRGBA(0, 0, 0, 0)
	c.Clear()

	c.SetColor(textColor)
	c.SetFontFace(fontFace)
	c.DrawStringAnchored(text, 0, 0, 0, 0.75)
	return c.Image()
}
