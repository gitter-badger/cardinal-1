package cards

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
