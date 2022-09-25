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

func DrawOver(background draw.Image, foreground image.Image) draw.Image {
	draw.Draw(background, background.Bounds(), foreground, foreground.Bounds().Min, draw.Over)
	return background
}

func GetRGBA(img image.Image) *image.RGBA {
	drawing := image.NewRGBA(img.Bounds())

	draw.Draw(drawing, drawing.Rect, img, img.Bounds().Min, draw.Over)
	return drawing
}