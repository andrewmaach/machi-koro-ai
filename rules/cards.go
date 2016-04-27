package rules

// Card is used to interact with cards from the game engine.
type Card interface {
	Roll() (int, int) // min, max
	Title() string
	BuildCost() int
	Kind() string
	Order() int
	Activate(*Game)
}

// IncomeCard is used for most blue cards and some green like convincience stroes.
type IncomeCard struct {
	title     string
	roll      int
	buildCost int
	anyone    bool
	kind      string
	income    int
}

// WheatField is a wheat field.
var WheatField = IncomeCard{
	title:     "Wheat Field",
	roll:      1,
	buildCost: 1,
	anyone:    true,
	kind:      "grain",
	income:    1,
}

// Ranch is a wheat field.
var Ranch = IncomeCard{
	title:     "Ranch",
	roll:      2,
	buildCost: 1,
	anyone:    true,
	kind:      "ranch",
	income:    1,
}

// FlowerOrchard is a wheat field.
var FlowerOrchard = IncomeCard{
	title:     "Flower Orchard",
	roll:      4,
	buildCost: 2,
	anyone:    true,
	kind:      "grain",
	income:    1,
}

// ConvenienceStore is a wheat field.
var ConvenienceStore = IncomeCard{
	title:     "ConvenienceS tore",
	roll:      4,
	buildCost: 2,
	anyone:    false,
	kind:      "grain",
	income:    3,
}

// Forest is a wheat field.
var Forest = IncomeCard{
	title:     "Forest",
	roll:      5,
	buildCost: 3,
	anyone:    true,
	kind:      "natural-resource",
	income:    1,
}

// Vineyard is a wheat field.
var Vineyard = IncomeCard{
	title:     "Vineyard",
	roll:      7,
	buildCost: 3,
	anyone:    true,
	kind:      "grain",
	income:    3,
}

// Mine is a wheat field.
var Mine = IncomeCard{
	title:     "Mine",
	roll:      9,
	buildCost: 6,
	anyone:    true,
	kind:      "natural-resource",
	income:    5,
}

// AppleOrchard is a wheat field.
var AppleOrchard = IncomeCard{
	title:     "Apple Orchard",
	roll:      10,
	buildCost: 3,
	anyone:    true,
	kind:      "grain",
	income:    3,
}
