package main

import (
	"juro-go/api"
	"juro-go/api/recipe"
	"juro-go/server"
)

func main() {
	app := server.Create()

	recipe.GetRecipeById(1)

	api.Setup(app)

	server.Listen(app)
}
