package environment

import "edoex/models"

var (
	Cards      map[int64]*models.Card  = make(map[int64]*models.Card)
	MetasIds   map[int64]*models.Meta  = make(map[int64]*models.Meta)
	MetasAlias map[string]*models.Meta = make(map[string]*models.Meta)

	Sets map[string]*models.Meta = make(map[string]*models.Meta)
)
