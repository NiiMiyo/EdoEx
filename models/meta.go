package models

import (
	"fmt"
	"strings"
)

// Any set or counter to be put on strings.conf
type Meta interface {
	// Line to put on strings.conf
	StringConfLine() string
	// Hexadecimal id
	HexId() string
}

type Set struct {
	Id    int64
	Name  string
	Alias string
}

func (self Set) HexId() string {
	hex := fmt.Sprintf("%x", self.Id)

	if difference := 4 - len(hex); difference > 0 {
		// if len(hex) < 4
		hex = strings.Repeat("0", difference) + hex
	} else if difference < 0 {
		// if len(hex) > 4
		hex = hex[-difference:]
	}

	return hex
}

func (self Set) StringConfLine() string {
	return fmt.Sprintf("!setname 0x%s %s", self.HexId(), self.Name)
}

type Counter struct {
	Id   int64
	Name string
}

func (self Counter) HexId() string {
	return fmt.Sprintf("%x", self.Id)
}

func (self Counter) StringConfLine() string {
	return fmt.Sprintf("!counter 0x%s %s", self.HexId(), self.Name)
}
