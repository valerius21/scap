package webserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Gin(receiverPort string) {
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
		ctx.Data(http.StatusOK, "application/json", msg)
	})

	r.GET("/empty", func(ctx *gin.Context) {
		msg, err := CreateHandler("gin", "empty", "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.Data(http.StatusOK, "application/json", msg)
	})

	r.GET("/sleep", func(ctx *gin.Context) {
		msg, err := CreateHandler("gin", "sleep", "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.Data(http.StatusOK, "application/json", msg)
	})

	r.POST("/image", func(ctx *gin.Context) {
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Image: could not get file",
			})
			return
		}

		args, err := ImageSaver(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Image: could not save file",
			})
			return
		}
		msg, err := CreateHandler("gin", "image", args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.Data(http.StatusOK, "application/json", msg)
	})

	r.Run(":" + receiverPort)
}
