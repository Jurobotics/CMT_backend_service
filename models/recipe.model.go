package models

type Tabler interface {
	TableName() string
}

type Recipe struct {
	Default
	Name        string `json:"name"`
	Description string `json:"description"`
	Preparation string `json:"preparation"`
	Image       string `json:"image"`
	Ingredients []Ingredients
}

func (Recipe) TableName() string {
	return "recipe"
}
