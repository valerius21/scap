package webserver

import (
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/sender"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valerius21/scap/pkg/fns"
)

func Gin(receiverHost, receiverPort string) {
	r := gin.Default()

	// Create a connection to the receiver
	s := sender.CreateSender(receiverHost, receiverPort)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing connection")
		}
	}(*s.C)

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
