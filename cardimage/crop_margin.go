package cardimage

import (
	"image"
	"image/draw"
)

var cropRect = image.Rect(23, 24, 716, 1036)

func CropMargin(img image.Image) image.Image {
	drawing := image.NewRGBA(cropRect)
	draw.Draw(drawing, cropRect, img, cropRect.Min, draw.Over)
	return drawing
}
