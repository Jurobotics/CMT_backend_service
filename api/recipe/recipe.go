package recipe

import "github.com/gofiber/fiber/v2"

func getRecipe(ctx *fiber.Ctx) error {
	return ctx.SendString("Recipe")
}
