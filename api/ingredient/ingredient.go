package ingredient

import (
	"juro-go/database"
	"juro-go/models"
)

type Ingredients struct {
	models.Ingredient
	AmountNeeded int `json:"amount_needed"`
}

func GetIngredientByRecipeId(id int) []Ingredients {
	var ingredients []Ingredients

	database.DB.Raw(`SELECT ingredient.name, ingredient.doser_size, ingredient.ID, i.amount as amount_needed FROM ingredient
         LEFT JOIN ingredients i on ingredient.id = i.ingredient_id
         LEFT JOIN recipe r on i.recipe_id = r.id WHERE r.id =?`, id).Scan(&ingredients)

	return ingredients
}
