package order

import "github.com/gofiber/fiber/v2"

func GetOrder(c *fiber.Ctx) error {
	return c.SendString("Order")
}
