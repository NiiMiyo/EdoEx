package cardimage

import "image"

type positions struct {
	Attribute, LinkArrows, SpellTrapType, Hologram image.Point
	Artwork, ArtworkPendulum                       image.Rectangle
}

var BuildPositions = positions{
	Attribute:     image.Point{604, 68},
	LinkArrows:    image.Point{57, 160},
	SpellTrapType: image.Point{365, 152},
	Hologram:      image.Point{660, 979},

	Artwork:         image.Rectangle{image.Point{107, 210}, image.Point{633, 736}},
	ArtworkPendulum: image.Rectangle{image.Point{68, 206}, image.Point{670, 654}},
}
