package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/valerius21/scap/pkg/fns"
	"net/http"
)

func Gin(receiverHost, receiverPort string) {
	r := gin.Default()

	r.GET("/math", func(ctx *gin.Context) {
		numberStr := ctx.Query("number")
		msg, err := CreateHandler("gin", "math", numberStr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, msg)
	})

	r.GET("/empty", func(ctx *gin.Context) {
		msg, err := CreateHandler("gin", "empty", "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, msg)
	})

	r.GET("/sleep", func(ctx *gin.Context) {
		msg, err := CreateHandler("gin", "sleep", "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, msg)
	})

	r.GET("/image", func(ctx *gin.Context) {
		fns.GenerateImageMetadataFn()
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"message": "not implemented",
		})
	})

	r.Run()
}
