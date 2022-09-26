package cardimage

import "image"

type positions struct {
	Attribute, LinkArrows, SpellTrapType image.Point
}

var BuildPositions = positions{
	Attribute:     image.Point{604, 68},
	LinkArrows:    image.Point{57, 160},
	SpellTrapType: image.Point{365, 152},
}
