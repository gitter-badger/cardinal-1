package cards

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
