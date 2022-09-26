package cardimage

import "image"

type positions struct {
	attribute, linkArrows image.Point
}

var BuildPositions = positions{
	attribute:  image.Point{604, 68},
	linkArrows: image.Point{57, 160},
}
