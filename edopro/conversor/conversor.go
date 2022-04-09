package conversor

// Values to convert scale to EDOPro's database values
var ScaleConversor struct {
	Left  int64
	Right int64
} = struct {
	Left  int64
	Right int64
}{
	Left:  0x1000000,
	Right: 0x10000,
}

// Maps from edoex yaml files' values to EDOPro's database values

var (
	Ruleset = map[string]int64{
		"ocg":        0x1,
		"tcg":        0x2,
		"animanga":   0x4,
		"illegal":    0x8,
		"videogame":  0x10,
		"custom":     0x20,
		"speed":      0x40,
		"prerelease": 0x100,
		"rush":       0x200,
		"legend":     0x400,
		"hidden":     0x1000,
	}

	Type = map[string]int64{
		"monster":        0x1,
		"spell":          0x2,
		"trap":           0x4,
		"normal":         0x10,
		"effect":         0x20,
		"fusion":         0x40,
		"ritual":         0x80,
		"spirit":         0x200,
		"union":          0x400,
		"gemini":         0x800,
		"tuner":          0x1000,
		"synchro":        0x2000,
		"token":          0x4000,
		"quickplay":      0x10000,
		"continuous":     0x20000,
		"equip":          0x40000,
		"field":          0x80000,
		"counter":        0x100000,
		"flip":           0x200000,
		"toon":           0x400000,
		"xyz":            0x800000,
		"pendulum":       0x1000000,
		"special_summon": 0x2000000,
		"link":           0x4000000,
		"skill":          0x8000000,
		"action":         0x10000000,
		"plus":           0x20000000,
		"minor":          0x40000000,
		"armor":          0x80000000,
	}

	LinkArrows = map[string]int64{
		"bottom_left":  0x1,
		"down":         0x2,
		"bottom_right": 0x4,
		"left":         0x8,
		"right":        0x20,
		"upper_left":   0x40,
		"top":          0x80,
		"upper_right":  0x100,
	}

	Race = map[string]int64{
		"warrior":       0x1,
		"spellcaster":   0x2,
		"fairy":         0x4,
		"fiend":         0x8,
		"zombie":        0x10,
		"machine":       0x20,
		"aqua":          0x40,
		"pyro":          0x80,
		"rock":          0x100,
		"winged_beast":  0x200,
		"plant":         0x400,
		"insect":        0x800,
		"thunder":       0x1000,
		"dragon":        0x2000,
		"beast":         0x4000,
		"beast_warrior": 0x8000,
		"dinosaur":      0x10000,
		"fish":          0x20000,
		"sea_serpent":   0x40000,
		"reptile":       0x80000,
		"psychic":       0x100000,
		"divine_beast":  0x200000,
		"creator_god":   0x400000,
		"wyrm":          0x800000,
		"cyberse":       0x1000000,
	}

	Attribute = map[string]int64{
		"earth":  0x1,
		"water":  0x2,
		"fire":   0x4,
		"wind":   0x8,
		"light":  0x10,
		"dark":   0x20,
		"divine": 0x40,
	}

	Category = map[string]int64{
		"destroy_monster":   0x1,
		"destroy_spelltrap": 0x2,
		"destroy_deck":      0x4,
		"destroy_hand":      0x8,
		"send_to_gy":        0x10,
		"send_to_hand":      0x20,
		"send_to_deck":      0x40,
		"banish":            0x80,
		"draw":              0x100,
		"search":            0x200,
		"change_atk_def":    0x400,
		"change_level_rank": 0x800,
		"position":          0x1000,
		"piercing":          0x2000,
		"direct_attack":     0x4000,
		"multi_attack":      0x8000,
		"negate_activation": 0x10000,
		"negate_effect":     0x20000,
		"damage_lp":         0x40000,
		"recover_lp":        0x80000,
		"special_summon":    0x100000,
		"non_effect":        0x200000,
		"token_related":     0x400000,
		"fusion_related":    0x800000,
		"ritual_related":    0x1000000,
		"synchro_related":   0x2000000,
		"xyz_related":       0x4000000,
		"link_related":      0x8000000,
		"counter_related":   0x10000000,
		"gamble":            0x20000000,
		"control":           0x40000000,
		"move_zones":        0x80000000,
	}
)
