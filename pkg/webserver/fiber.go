package webserver

import (
	"github.com/valerius21/scap/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Fiber is a webserver implementation using the Fiber framework on top of fastHTTP
func Fiber(receiverHost, receiverPort string) {
	// Create a new Fiber instance
	app := fiber.New()

	// Define the routes and their corresponding handlers
	app.Get("/image", func(c *fiber.Ctx) error {
		defer utils.TimeTrack(time.Now(), "Fiber:ImageHandler")

		return c.SendString("imageHandler: not implemented")
	})
	app.Get("/empty", func(ctx *fiber.Ctx) error {
		msg, err := CreateHandler("fiber", "empty", "")
		if err != nil {
			return err
		}
		return ctx.Send(msg)
	})

	app.Get("/math", func(c *fiber.Ctx) error {
		n := c.Query("number")
		msg, err := CreateHandler("fiber", "math", n)
		if err != nil {
			return err
		}
		return c.Send(msg)
	})

	app.Get("/sleep", func(c *fiber.Ctx) error {
		msg, err := CreateHandler("fiber", "sleep", "")
		if err != nil {
			return err
		}
		return c.Send(msg)
	})

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
		return
	}
}
