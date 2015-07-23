package cards

// Card allows us to store all kinds of cards regardless of game in a Collection's Cards array
// so all cards implement the Card interface.
type Card interface {
	GetID() string
	GetName() string
}

// Collection is a collection of cards. It has an owner, name, game, slice of cards, and whether it is the main collection or not.
type MagicCollection struct {
	Name   string      `json:"name"`
	Game   string      `json:"game"`
	IsMain bool        `json:"ismain"`
	Cards  []MagicCard `json:"cards"`
}
