package ingredient

import (
	"juro-go/database"
	"juro-go/models"
)

func GetIngredientByRecipeId(id int) []models.Ingredient {
	var ingredients []models.Ingredient

	database.DB.Raw(`SELECT ingredient.name, ingredient.amount, i.amount as pumps, s.arduino FROM ingredient
         LEFT JOIN ingredients i on ingredient.id = i.ingredient_id
         LEFT JOIN recipe r on i.recipe_id = r.id
        LEFT OUTER JOIN servos s on s.id = ingredient.servo_id
WHERE r.id =?`, id).Scan(&ingredients)

	return ingredients
}
