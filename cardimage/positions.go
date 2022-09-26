package cardimage

import "image"

type positions struct {
	Attribute, LinkArrows, SpellTrapType, Hologram image.Point
	ArtworkBox, ArtworkPendulumBox, NameBox        image.Rectangle
}

var BuildPositions = positions{
	Attribute:     image.Point{604, 68},
	LinkArrows:    image.Point{57, 160},
	SpellTrapType: image.Point{365, 152},
	Hologram:      image.Point{660, 979},

	ArtworkBox:         image.Rect(107, 210, 633, 736),
	ArtworkPendulumBox: image.Rect(68, 206, 670, 654),
	NameBox:            image.Rect(76, 69, 592, 133),
}
