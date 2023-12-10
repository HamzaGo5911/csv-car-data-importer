package models

import "github.com/fatih/structs"

// Car holds information for cars
type Car struct {
	ID           string `json:"id" structs:"id"`
	Name         string `json:"name" structs:"name"`
	Year         string `json:"year" structs:"year"`
	SellingPrice string `json:"selling_price" structs:"selling_price"`
	Transmission string `json:"transmission" structs:"transmission"`
}

// Map converts structs to a map representation
func (c *Car) Map() map[string]interface{} {
	return structs.Map(c)
}

// Names returns the field names of Car model.
func (c *Car) Names() []string {
	fields := structs.Fields(c)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
