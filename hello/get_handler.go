package hello

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nonbody/sample-app/print"
)

func NewGetHandler(f *fiber.App, p print.PrintFunc) {

	f.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋! " + p("xxx"))
	})

}
