package recipe

import (
	"errors"
	"juro-go/api/ingredient"
	"juro-go/models"
	"juro-go/pkg"
	"strconv"

	"juro-go/database"

	"github.com/gofiber/fiber/v2"
)

type RecipeData struct {
	Recipe      models.Recipe            `json:"recipe"`
	Ingredients []ingredient.Ingredients `json:"ingredients"`
}

func GetRecipeById(id int) RecipeData {
	var res RecipeData
	database.DB.Model(&models.Recipe{}).Where("id=?", id).First(&res.Recipe)

	res.Ingredients = ingredient.GetIngredientByRecipeId(id)

	return res
}

func getRecipes(ctx *fiber.Ctx) error {
	var recipes []models.Recipe
	result := database.DB.Find(&recipes)
	if result.RowsAffected == 0 {
		return ctx.JSON(pkg.ErrorResponse(errors.New("no recipes found")))
	}
	return ctx.JSON(pkg.SuccessResponse(recipes))
}

func getRecipeById(ctx *fiber.Ctx) error {
	recipeId := ctx.Params("id")

	id, _ := strconv.Atoi(recipeId)

	return ctx.JSON(pkg.SuccessResponse(GetRecipeById(id)))
}

func createRecipe(ctx *fiber.Ctx) error {

	recipe := new(models.Recipe)

	if err := ctx.BodyParser(recipe); err != nil {
		return ctx.JSON(pkg.ErrorResponse(err))
	}

	database.DB.Create(&recipe)

	return ctx.JSON(pkg.SuccessResponse("recipe created"))
}
