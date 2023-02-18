package order

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/:recipeId", Order)
}
