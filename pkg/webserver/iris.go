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

	app.Get("/image", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotImplemented)
		ctx.Text("Image handler: not implemented")
	})

	app.Listen(":" + receiverPort)
}
