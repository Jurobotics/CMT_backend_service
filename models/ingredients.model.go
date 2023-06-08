package models

type Ingredients struct {
	Default
	Amount       float32    `json:"amount"`
	AmountSolid  string     `json:"amountSolid"`
	IngredientID int        `json:"ingredientId"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID;references:ID"`
	RecipeID     int        `json:"recipeId"`
	Recipe       Recipe     `gorm:"foreignKey:RecipeID;references:ID"`
}
