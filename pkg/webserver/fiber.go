package webserver

import (
	"fmt"
	"github.com/valerius21/scap/pkg/sender"
	"github.com/valerius21/scap/pkg/utils"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/fns"
)

// Fiber is a webserver implementation using the Fiber framework on top of fastHTTP
func Fiber(receiverHost, receiverPort string) {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a connection to the receiver
	s := sender.CreateSender(receiverHost, receiverPort)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing connection")
		}
	}(*s.C)

	// Define the routes and their corresponding handlers
	app.Get("/image", func(c *fiber.Ctx) error {
		defer utils.TimeTrack(time.Now(), "Fiber:ImageHandler")

		return c.SendString("imageHandler: not implemented")
	})
	app.Get("/empty", func(c *fiber.Ctx) error {
		defer utils.TimeTrack(time.Now(), "Fiber:EmptyHandler")
		_, err := s.Send("empty:")
		if err != nil {
			return err
		}
		return c.SendString("EmptyHandler: Executed")
	})
	app.Get("/math", func(c *fiber.Ctx) error {
		defer utils.TimeTrack(time.Now(), "Fiber:MathHandler")
		number := c.QueryInt("number")
		result := fns.MathFn(number)
		return c.SendString(fmt.Sprintf("MathHandler: %.5f", result))
	})
	app.Get("/sleep", func(c *fiber.Ctx) error {
		defer utils.TimeTrack(time.Now(), "Fiber:SleepHandler")
		fns.SleeperFn(1 * time.Second)
		return c.SendString("sleepHandler: slept for 1 second")
	})

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
		return
	}
}
