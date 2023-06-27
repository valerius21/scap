package webserver

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valerius21/scap/pkg/fns"
)

func Fiber() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define the routes and their corresponding handlers
	app.Get("/image", func(c *fiber.Ctx) error {
		return c.SendString("imageHandler: not implemented")
	})

	app.Get("/emtpy", func(c *fiber.Ctx) error {
		fns.EmptyFn()
		return c.SendString("EmptyHandler: Executed")
	})
	app.Get("/math", func(c *fiber.Ctx) error {
		number := c.QueryInt("number")
		result := fns.MathFn(number)
		return c.SendString(fmt.Sprintf("MathHandler: %.2f", result))
	})
	app.Get("/sleep", func(c *fiber.Ctx) error {
		fns.SleeperFn(1 * time.Second)
		return c.SendString("sleepHandler: slept for 1 second")
	})

	// Start the server
	app.Listen(":8080")
}
