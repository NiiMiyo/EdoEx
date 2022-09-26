package cardimage

import "image"

type positions struct {
	Attribute, LinkArrows, SpellTrapType, Hologram image.Point
}

var BuildPositions = positions{
	Attribute:     image.Point{604, 68},
	LinkArrows:    image.Point{57, 160},
	SpellTrapType: image.Point{365, 152},
	Hologram:      image.Point{660, 979},
}
