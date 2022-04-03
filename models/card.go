package models

import (
	"edoex/parser"
	"edoex/utils/maputils"
	"edoex/utils/sliceutils"
	"strconv"
)

type Card struct {
	// todo: Support different scales for each side
	// todo: Allow setcodes of sets that are not on the expansion
	// todo: lowercase ruleset, type, race, attribute, category on parse

	Id          int64    // Card code
	Name        string   // Card name
	Description string   // Card text (Flavor text or effect)
	CardType    string   // Monster, Spell or Trap
	SubTypes    []string // Effect, Pendulum, Continuous and this kind of stuff

	Atk                 int64    // Monster's attack points
	Def                 int64    // Monster's defense points
	Level               int64    // Monster's level, rank ou link rating
	Race                []string // Monster's type (Warrior, Plant...)
	Attribute           []string // Monster's attribute
	PendulumDescription string   // Pendulum text box
	LinkArrows          []string // Link arrow's directions
	Scale               int64    // Pendulum scale

	Ruleset  []string // OCG, TCG, Anime...
	Alias    int64    // Alias code
	Sets     []Set
	Category []string // Categories for search
	Strings  []string
}

// Card data following EDOPro's database structure to simplify database insertion
type CardDb struct {
	Id        int64
	Ot        int64
	Alias     int64
	Setcode   int64
	Type      int64
	Atk       int64
	Def       int64
	Level     int64
	Race      int64
	Attribute int64
	Category  int64

	Name    string
	Desc    string
	Strings [16]string
}

// Returns same card as a CardDb
func (self *Card) ToDb() CardDb {
	return CardDb{
		Id:        self.Id,
		Ot:        self.getOt(),
		Alias:     self.Alias,
		Setcode:   self.getSetcode(),
		Type:      self.getType(),
		Atk:       self.Atk,
		Def:       self.getDef(),
		Level:     self.getLevel(),
		Race:      self.getRace(),
		Attribute: self.getAttribute(),
		Category:  self.getCategory(),
		Name:      self.Name,
		Desc:      self.getDesc(),
		Strings:   self.getStrings(),
	}
}

func (self *Card) hasSubType(sub string) bool {
	return sliceutils.Contains(self.SubTypes, sub)
}

func (self *Card) getOt() int64 {
	ot := int64(0)

	for _, r := range self.Ruleset {
		ot += maputils.GetValueOrDefault(parser.Maps.Ruleset, r, 0)
	}

	return ot
}

func (self *Card) getSetcode() int64 {
	setcode := ""

	for _, s := range self.Sets {
		setcode += s.HexId()
	}

	code, _ := strconv.ParseInt(setcode, 16, 64)
	return code
}

func (self *Card) getType() int64 {
	_type := int64(parser.Maps.Type[self.CardType])

	for _, t := range self.SubTypes {
		_type += maputils.GetValueOrDefault(parser.Maps.Type, t, 0)
	}

	return _type
}

func (self *Card) getDef() int64 {
	if self.hasSubType("link") {
		arrows := int64(0)

		for _, a := range self.LinkArrows {
			arrows += maputils.GetValueOrDefault(parser.Maps.LinkArrows, a, 0)
		}

		return arrows
	}

	return self.Def
}

func (self *Card) getLevel() int64 {
	level := self.Level

	if self.hasSubType("pendulum") {
		level += (self.Scale * parser.ScaleConversor.Left) + (self.Scale * parser.ScaleConversor.Right)
	}

	return level
}

func (self *Card) getRace() int64 {
	race := int64(0)

	for _, r := range self.Race {
		race += maputils.GetValueOrDefault(parser.Maps.Race, r, 0)
	}

	return race
}

func (self *Card) getAttribute() int64 {
	attribute := int64(0)

	for _, a := range self.Attribute {
		attribute += maputils.GetValueOrDefault(parser.Maps.Attribute, a, 0)
	}

	return attribute
}

func (self *Card) getCategory() int64 {
	category := int64(0)

	for _, c := range self.Category {
		category += maputils.GetValueOrDefault(parser.Maps.Category, c, 0)
	}

	return category
}

func (self *Card) getDesc() string {
	if self.hasSubType("pendulum") {
		var descType string
		if self.hasSubType("effect") {
			descType = "[ Monster Effect ]\n"
		} else {
			descType = "[ Flavor Text ]\n"
		}

		return "[ Pendulum Effect ]\n" + self.PendulumDescription +
			"\n----------------------------------------\n" + descType + self.Description
	}

	return self.Description
}

func (self *Card) getStrings() [16]string {
	strings := [16]string{}

	for i, s := range self.Strings {
		if i >= 16 {
			break
		}

		strings[i] = s
	}

	return strings
}