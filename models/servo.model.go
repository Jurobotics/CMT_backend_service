package models

type Servo struct {
	Default
	Arduino    int `json:"arduino"`
	Ingredient Ingredient
}
