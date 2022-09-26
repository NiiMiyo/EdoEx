package cardimage

import "image"

type positions struct {
	attribute image.Point
}

var BuildPositions positions

func init() {
	BuildPositions.attribute = image.Pt(604, 68)
}
