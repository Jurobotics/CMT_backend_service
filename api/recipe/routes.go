package recipe

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/", getRecipes)
	route.Get("/:id", getRecipeById)
	route.Post("/", createRecipe)
}
