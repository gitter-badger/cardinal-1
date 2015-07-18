package cards

// Legality is for Magic Cards there is a nested JSON object which contains the Legality information about each card.
type Legality struct {
	Modern           string `json:"modern"`
	Legacy           string `json:"legacy"`
	Vintage          string `json:"vintage"`
	Freeform         string `json:"freeform"`
	Prismatic        string `json:"prismatic"`
	TribalWarsLegacy string `json:"tribalwarslegacy"`
	Singleton        string `json:"singleton"`
	Commander        string `json:"commander"`
}

// ForeignName is also for Magic Cards there is a nested JSON object for each foreign name of a given card.
type ForeignName struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

// MagicCard represents a Magic: The Gathering card and implements the Card Interface
type MagicCard struct {
	ID           string        `json:"id"`
	Layout       string        `json:"layout"`
	Name         string        `json:"name"`
	ManaCost     string        `json:"manaCost"`
	Cmc          int           `json:"cmc"`
	Colors       []string      `json:"colors"`
	Type         string        `json:"type"`
	Types        []string      `json:"types"`
	SubTypes     []string      `json:"subTypes"`
	Power        int           `json:"power"`
	Toughness    int           `json:"toughness"`
	Text         string        `json:"text"`
	ForeignNames []ForeignName `json:"foreignNames"`
	Printings    []string      `json:"printings"`
	Legalities   Legality      `json:"legalities"`
	ImageNames   []string      `json:"imageNames"`
}

// GetID implements the Card Interface
func (m MagicCard) GetID() string {
	return m.ID
}

// GetName implements the Card Interface
func (m MagicCard) GetName() string {
	return m.Name
}
