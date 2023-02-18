package models

type Ingredient struct {
	Default
	Name      string `json:"name"`
	DoserSize int    `json:"doser_size"` // Amount of the doser in the Si unit
	Recipes   []Ingredients
}

func (Ingredient) TableName() string {
	return "ingredient"
}
