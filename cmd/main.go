package main

import (
	"flag"
	"net"
	"net/rpc"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/dto"
	"github.com/valerius21/scap/pkg/rpc_services"
	"github.com/valerius21/scap/pkg/webserver"
)

func main() {
	if _, err := os.Stat("/tmp/scap"); os.IsNotExist(err) {
		err := os.Mkdir("/tmp/scap", 0o777)
		if err != nil {
			log.Error().Err(err).Msg("Error when creating the tmp directory")
			return
		}
	}
	modePtr := flag.Bool("is-server", false, "if true, run the webserver in http mode,"+
		" otherwise operator mode")
	webServerPtr := flag.String("webserver", "net", "the webserver to run")
	portPtr := flag.String("port", "3000", "the port to run the http on")

	flag.Parse()

	if *modePtr {
		log.Info().Msg("Running in server mode")
		handlerService := new(rpc_services.HandlerService)
		rpc.Register(handlerService)

		listener, err := net.Listen("tcp", "127.0.0.1:1234")
		if err != nil {
			panic(err)
		}

		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			go rpc.ServeConn(conn)
		}
	} else {
		log.Info().Msg("Running in client mode")
		var response dto.Message

		log.Info().Msgf("The server time is: %v", response)
		switch *webServerPtr {
		case "fiber":
			webserver.Fiber(*portPtr)
			return
		case "echo":
			webserver.Echo(*portPtr)
			return
		case "gin":
			webserver.Gin(*portPtr)
			return
		case "iris":
			webserver.Iris(*portPtr)
			return
		default:
			webserver.NetHttp(*portPtr)
			return
		}

	}
}
