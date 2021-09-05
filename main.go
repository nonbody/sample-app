package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nonbody/sample-app/hello"
	px "github.com/nonbody/sample-app/print"
	"github.com/nonbody/sample-app/world"
)

func main() {

	app := fiber.New()

	// init db

	// init redis

	pnt := px.NewPrintFunc()

	hello.NewGetHandler(app, pnt)

	wget := world.NewGetHandler(pnt)

	// app.Get("/hello/get", hget)
	app.Post("/world/get", wget)

	app.Listen(":3000")

}
