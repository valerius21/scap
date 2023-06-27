package webserver

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/valerius21/scap/pkg/fns"
)

func Iris() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/math", func(ctx iris.Context) {
		nubmerStr := ctx.Request().URL.Query().Get("number")
		number, err := strconv.Atoi(nubmerStr)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Text("Invalid number")
			return
		}
		result := fns.MathFn(number)
		ctx.Text(fmt.Sprintf("Math handler: %.2f", result))
	})

	app.Get("/sleep", func(ctx iris.Context) {
		fns.SleeperFn(1 * time.Second)
		ctx.Text("Sleep handler: slept for 1 second")
	})

	app.Get("/empty", func(ctx iris.Context) {
		fns.EmptyFn()
		ctx.StatusCode(iris.StatusNoContent)
	})

	app.Get("/image", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotImplemented)
		ctx.Text("Image handler: not implemented")
	})

	app.Listen(":8080")
}
