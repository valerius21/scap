package webserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valerius21/scap/pkg/fns"
)

func Gin() {
	r := gin.Default()

	r.GET("/math", func(ctx *gin.Context) {
		numberStr := ctx.Query("number")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid number",
			})
			return
		}

		result := fns.MathFn(number)

		ctx.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.GET("/empty", func(ctx *gin.Context) {
		fns.EmptyFn()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/sleep", func(ctx *gin.Context) {
		fns.SleeperFn(1 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/image", func(ctx *gin.Context) {
		fns.GenerateImageMetadataFn()
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"message": "not implemented",
		})
	})

	r.Run()
}
