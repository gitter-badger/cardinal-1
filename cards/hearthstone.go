package cards

// HearthStoneCard represents a given Hearth Stone Card
type HearthStoneCard struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Cost         int      `json:"cost"`
	Type         string   `json:"type"`
	Rarity       string   `json:"rarity"`
	Faction      string   `json:"faction"`
	Text         string   `json:"text"`
	Mechanics    []string `json:"mechanics"`
	Flavor       string   `json:"flavor"`
	Artist       string   `json:"artist"`
	Attack       int      `json:"attack"`
	Health       int      `json:"health"`
	Collectible  bool     `json:"collectible"`
	Elite        bool     `json:"elite"`
	PlayerClass  string   `json:"playerClass"`
	InPlayText   string   `json:"inPlayText"`
	Durability   int      `json:"durability"`
	HowToGet     string   `json:"howToGet"`
	HowToGetGold string   `json:"hotToGetGold"`
}

// GetID implements the Card Interface
func (h HearthStoneCard) GetID() string {
	return h.ID
}

// GetName implements the Card Interface
func (h HearthStoneCard) GetName() string {
	return h.Name
}
