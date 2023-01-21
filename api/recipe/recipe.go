package recipe

import (
	"juro-go/models"
	"juro-go/pkg"

	"juro-go/database"

	"github.com/gofiber/fiber/v2"
)

type RecipeBody struct {
}

func getRecipe(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}

func getRecipeById(ctx *fiber.Ctx) error {
	recipeId := ctx.Params("id")

	var recipe models.Recipe

	result := database.DB.Where("id = ?", recipeId).First(&recipe)

	if result.RowsAffected == 0 {
		return ctx.SendStatus(404)
	}

	return ctx.JSON(&recipe)
}

func createRecipe(ctx *fiber.Ctx) error {

	recipe := new(models.Recipe)

	if err := ctx.BodyParser(recipe); err != nil {
		return ctx.JSON(pkg.ErrorResponse(err))
	}

	database.DB.Create(&recipe)

	return ctx.JSON(pkg.SuccessResponse("recipe created"))
}
