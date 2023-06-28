package webserver

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/sender"
	"net"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/valerius21/scap/pkg/fns"
)

func Echo(receiverHost, receiverPort string) {
	e := echo.New()

	// Create a connection to the receiver
	s := sender.CreateSender(receiverHost, receiverPort)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing connection")
		}
	}(*s.C)

	e.GET("/math", func(c echo.Context) error {
		numberStr := c.QueryParam("number")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return c.String(400, "Invalid number")
		}
		result := fns.MathFn(number)
		return c.String(200, fmt.Sprintf("Result: %.2f", result))
	})

	e.GET("/empty", func(c echo.Context) error {
		fns.EmptyFn()
		return c.String(200, "Empty")
	})

	e.GET("/image", func(c echo.Context) error {
		fns.GenerateImageMetadataFn()
		return c.String(500, "Image: not implemented")
	})

	e.GET("/sleep", func(c echo.Context) error {
		fns.SleeperFn(1 * time.Second)
		return c.String(200, "Sleep: 1s")
	})
}
