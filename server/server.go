package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Create() *fiber.App {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OKAY")
	})

	return app
}

func Listen(app *fiber.App) error {
	serverPort := os.Getenv("SERVER_PORT")

	if serverPort == "" {
		serverPort = "3000"
	}

	return app.Listen(fmt.Sprintf(":%s", serverPort))
}
