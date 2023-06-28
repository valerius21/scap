package webserver

import (
	"github.com/kataras/iris/v12"
)

func Iris(receiverPort string) {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/math", func(ctx iris.Context) {
		number := ctx.Request().URL.Query().Get("number")
		msg, err := CreateHandler("iris", "math", number)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Text(err.Error())
			return
		}
		ctx.Write(msg)
	})

	app.Get("/sleep", func(ctx iris.Context) {
		msg, err := CreateHandler("iris", "sleep", "")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Text(err.Error())
			return
		}
		ctx.Write(msg)
	})

	app.Get("/empty", func(ctx iris.Context) {
		msg, err := CreateHandler("iris", "empty", "")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Text(err.Error())
			return
		}
		ctx.Write(msg)
	})

	app.Post("/image", func(ctx iris.Context) {
		_, header, err := ctx.FormFile("image")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Text("Image: could not get file")
			return
		}

		args, err := ImageSaver(header)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Text("Image: could not save file")
			return
		}

		msg, err := CreateHandler("iris", "image", args)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Text(err.Error())
			return
		}

		ctx.Write(msg)
	})

	app.Listen(":" + receiverPort)
}
