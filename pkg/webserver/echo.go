package webserver

import (
	"github.com/labstack/echo/v4"
)

func Echo(receiverPort string) {
	e := echo.New()

	e.GET("/math", func(c echo.Context) error {
		numberStr := c.QueryParam("number")
		msg, err := CreateHandler("echo", "math", numberStr)
		if err != nil {
			return err
		}
		return c.JSONBlob(200, msg)
	})

	e.GET("/empty", func(c echo.Context) error {
		msg, err := CreateHandler("echo", "empty", "")
		if err != nil {
			return err
		}
		return c.JSONBlob(200, msg)
	})

	e.POST("/image", func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(500, "Image: not implemented")
		}
		args, err := ImageSaver(file)
		if err != nil {
			return c.JSON(500, "Image: not implemented")
		}

		msg, err := CreateHandler("echo", "image", args)
		if err != nil {
			return c.JSON(500, "Image: not implemented")
		}
		return c.JSONBlob(200, msg)
	})

	e.GET("/sleep", func(c echo.Context) error {
		msg, err := CreateHandler("echo", "sleep", "")
		if err != nil {
			return err
		}
		return c.JSONBlob(200, msg)
	})

	e.Logger.Fatal(e.Start(":" + receiverPort))
}
