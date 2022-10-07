package models

import (
	"fmt"
	"strings"
)

type MetaType uint

const (
	MetaTypeSet MetaType = 1 << iota
	MetaTypeCounter
	MetaTypeAlias
)

// Any set, meta or alias to be put on strings.conf
type Meta struct {
	Id    int64
	Type  MetaType
	Name  string
	Alias string
}

func (meta Meta) HexId() string {
	if meta.Type == MetaTypeSet {
		hex := fmt.Sprintf("%x", meta.Id)

		if difference := 4 - len(hex); difference > 0 {
			// if len(hex) < 4
			hex = strings.Repeat("0", difference) + hex
		} else if difference < 0 {
			// if len(hex) > 4
			hex = hex[-difference:]
		}

		return hex
	} else {
		return fmt.Sprintf("%x", meta.Id)
	}
}

func (meta Meta) StringConfLine() string {
	var confType string
	if meta.Type == MetaTypeSet && meta.AliasOrName() != "" {
		confType = "setname"
	} else if meta.Type == MetaTypeCounter {
		confType = "counter"
	} else {
		return ""
	}

	return fmt.Sprintf("!%s 0x%s %s", confType, meta.HexId(), meta.Name)
}

func (meta Meta) AliasOrName() string {
	if meta.Alias == "" {
		return meta.Name
	} else {
		return meta.Alias
	}
}
