package recipe

import (
	"juro-go/database"
	"juro-go/models"

	"github.com/gofiber/fiber/v2"
)

func getRecipe(ctx *fiber.Ctx) error {
	return ctx.SendString("Recipe")
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
