package cards

// MagicCard represents a Magic: The Gathering card and implements the Card Interface
type MagicCard struct {
	ID           string              `json:"id"`
	Layout       string              `json:"layout"`
	Name         string              `json:"name"`
	ManaCost     string              `json:"manaCost"`
	Cmc          int                 `json:"cmc"`
	Colors       []string            `json:"colors"`
	Type         string              `json:"type"`
	Types        []string            `json:"types"`
	SubTypes     []string            `json:"subTypes"`
	Power        int                 `json:"power"`
	Toughness    int                 `json:"toughness"`
	Text         string              `json:"text"`
	ForeignNames []map[string]string `json:"foreignNames"`
	Printings    []string            `json:"printings"`
	Legalities   map[string]string   `json:"legalities"`
	ImageNames   []string            `json:"imageNames"`
}

// GetID implements the Card Interface
func (m MagicCard) GetID() string {
	return m.ID
}

// GetName implements the Card Interface
func (m MagicCard) GetName() string {
	return m.Name
}

// String implements the Stringer Interface
func (m MagicCard) String() string {
	var s = ""
	s += "Layout: " + m.Layout + "\n"
	s += "Name: " + m.Name + "\n"
	s += "ManaCost: " + m.ManaCost + "\n"
	s += "CMC: " + string(m.Cmc) + "\n"
	s += "Colors: "
	for color := range m.Colors {
		s += m.Colors[color] + ", "
	}
	s += "\n"
	s += "Type: " + m.Type + "\n"
	s += "Types: "
	for typ := range m.Types {
		s += m.Types[typ] + ", "
	}
	s += "\n"
	s += "SubTypes: "
	for stype := range m.SubTypes {
		s += m.SubTypes[stype] + ", "
	}
	s += "\n"
	s += "Power: " + string(m.Power) + "\n"
	s += "Toughness: " + string(m.Toughness) + "\n"
	s += "Text: " + m.Text + "\n"
	s += "ForeignNames: \n"
	for fn := range m.ForeignNames {
		for k, v := range m.ForeignNames[fn] {
			s += "\t" + k + " " + v + "\n"
		}
	}
	s += "Printings: "
	for printing := range m.Printings {
		s += m.Printings[printing] + ", "
	}
	s += "\n"
	s += "Legalities: \n"
	for format, legal := range m.Legalities {
		s += "\t" + format + " " + legal + "\n"
	}
	return s
}
