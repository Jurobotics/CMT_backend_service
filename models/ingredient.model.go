package models

type Ingredient struct {
	Default
	Name    string `json:"name"`
	Amount  uint   `json:"amount"` // Amount of the doser in the Si unit
	ServoId uint   `json:"servoId"`
	Recipes []Ingredients
}

func (Ingredient) TableName() string {
	return "ingredient"
}
