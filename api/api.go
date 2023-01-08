package api

import (
	"juro-go/api/order"
	"juro-go/api/recipe"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API")
	})
	// order.Routes(v1)
	orderGroup := v1.Group("/order")
	order.Routes(orderGroup)

	recipeGroup := v1.Group("/recipe")
	recipe.Routes(recipeGroup)
}
