package imagesutils

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func LoadImageFromPath(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func DrawOver(background draw.Image, foreground image.Image) {
	DrawAt(background, foreground, foreground.Bounds().Min)
}

func GetRGBA(img image.Image) *image.RGBA {
	drawing := image.NewRGBA(img.Bounds())

	draw.Draw(drawing, drawing.Rect, img, img.Bounds().Min, draw.Over)
	return drawing
}

func DrawAt(dst draw.Image, src image.Image, at image.Point) {
	b := src.Bounds()
	at = at.Sub(b.Min)
	r := b.Add(at)

	draw.Over.Draw(dst, r, src, b.Min)
}

func Crop(img image.Image, rect image.Rectangle) draw.Image {
	drawing := image.NewRGBA(rect)
	draw.Draw(drawing, rect, img, rect.Min, draw.Over)
	return drawing
}
