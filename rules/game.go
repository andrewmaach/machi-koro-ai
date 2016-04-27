package rules

// Player represents a unique player with cards, money.
type Player struct {
	Title string
	Cards []Card
	Coins int
}

// Deck represents the cards that have not yet been shown.
type Deck []Card

// Market represents the cards available to choose from.
type Market []Card

// Game represents the current players and cards on the table.
type Game struct {
	Players []Player
	Deck    Deck
	Market  Market
}
