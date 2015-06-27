package main

// Card allows us to store all kinds of cards regardless of game in a Collection's Cards array
// so all cards implement the Card interface.
type Card interface {
	GetID() string
	GetName() string
}

// User is to hold the user information in our DB
type User struct {
	Username    string
	Password    []byte `json:"-"`
	DashItems   []DashItem
	Collections []Collection
}

// Legality is for Magic Cards there is a nested JSON object which contains the Legality information about each card.
type Legality struct {
	Modern           string
	Legacy           string
	Vintage          string
	Freeform         string
	Prismatic        string
	TribalWarsLegacy string
	Singleton        string
	Commander        string
}

// ForeignName is also for Magic Cards there is a nested JSON object for each foreign name of a given card.
type ForeignName struct {
	Language string
	Name     string
}

// HearthStoneCard represents a given Hearth Stone Card
type HearthStoneCard struct {
	ID           string
	Name         string
	Cost         int
	Type         string
	Rarity       string
	Faction      string
	Text         string
	Mechanics    []string
	Flavor       string
	Artist       string
	Attack       int
	Health       int
	Collectible  bool
	Elite        bool
	PlayerClass  string
	InPlayText   string
	Durability   int
	HowToGet     string
	HowToGetGold string
}

// GetID implements the Card Interface
func (h HearthStoneCard) GetID() string {
	return h.ID
}

// GetName implements the Card Interface
func (h HearthStoneCard) GetName() string {
	return h.Name
}

// MagicCard represents a Magic: The Gathering card and implements the Card Interface
type MagicCard struct {
	ID           string
	Layout       string
	Name         string
	ManaCost     string
	Cmc          int
	Colors       []string
	Type         string
	Types        []string
	SubTypes     []string
	Power        int
	Toughness    int
	Text         string
	ForeignNames []ForeignName
	Printings    []string
	Legalities   Legality
	ImageNames   []string
}

// GetID implements the Card Interface
func (m MagicCard) GetID() string {
	return m.ID
}

// GetName implements the Card Interface
func (m MagicCard) GetName() string {
	return m.Name
}

// Collection is a collection of cards. It has an owner, name, game, slice of cards, and whether it is the main collection or not.
type Collection struct {
	Name   string
	Game   string
	IsMain bool
	Owner  User
	Cards  []Card
}

// DashItem is the basic form of a "Dash Card" which displays various information to the user.
type DashItem struct {
	Img     string
	Title   string
	Content string
}
