package models

type Tabler interface {
	TableName() string
}

type Recipe struct {
	Default
	Name        string `json:"name"`
	Image       string `json:"image"`
	Alc         bool
	Ingredients []Ingredients
}

func (Recipe) TableName() string {
	return "recipe"
}
