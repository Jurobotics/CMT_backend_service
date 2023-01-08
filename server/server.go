package server

import (
	"fmt"
	"juro-go/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Create() *fiber.App {
	database.SetupDatabase("./data/test.db")

	app := fiber.New()

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
