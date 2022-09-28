package cardimage

import "image"

type positions struct {
	Attribute, LinkArrows, SpellTrapType, Hologram, Stars, Defense,
	LinkRating image.Point
	ArtworkBox, ArtworkPendulumBox, NameBox, AbilitiesBox,
	TextMonsterBox, TextSpellTrapBox image.Rectangle
}

var BuildPositions = positions{
	Attribute:     image.Point{604, 68},
	LinkArrows:    image.Point{57, 160},
	SpellTrapType: image.Point{365, 152},
	Hologram:      image.Point{660, 979},
	Stars:         image.Point{599, 146},
	Defense:       image.Point{662, 940},
	LinkRating:    image.Point{662, 945},

	ArtworkBox:         image.Rect(107, 210, 633, 736),
	ArtworkPendulumBox: image.Rect(68, 206, 670, 654),
	NameBox:            image.Rect(76, 69, 592, 133),
	AbilitiesBox:       image.Rect(77, 793, 662, 830),
	TextMonsterBox:     image.Rect(75, 825, 662, 937),
	TextSpellTrapBox:   image.Rect(75, 787, 662, 970),
}
