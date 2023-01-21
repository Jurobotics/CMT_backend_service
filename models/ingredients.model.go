package models

type Ingredients struct {
	Default
	Amount       uint       `json:"amount"`
	IngredientID uint       `json:"ingredientId"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID;references:ID"`
	RecipeID     uint       `json:"recipeId"`
	Recipe       Recipe     `gorm:"foreignKey:RecipeID;references:ID"`
}
