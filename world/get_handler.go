package world

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nonbody/sample-app/print"
)

func NewGetHandler(p print.PrintFunc) fiber.Handler {

	fmt.Println("init....")

	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋! " + p("xxx"))
	}
}
