package main

import (
	"juro-go/api"
	"juro-go/server"
)

func main() {
	app := server.Create()

	api.Setup(app)

	server.Listen(app)
}
