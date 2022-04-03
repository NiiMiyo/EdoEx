package models

type Card struct {
	// todo: Support different scales for each side
	// todo: Allow setcodes of sets that are not on the expansion

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
