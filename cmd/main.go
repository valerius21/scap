package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/nsq"
	"github.com/valerius21/scap/pkg/webserver"
)

func main() {
	webServerPtr := flag.String("webserver", "fiber", "the webserver to run")
	modePtr := flag.Bool("http", false, "if true, run the webserver in http mode,"+
		" otherwise operator mode")
	hostPtr := flag.String("host", "localhost", "the host to run the tcp on")
	portPtr := flag.String("port", "3000", "the port to run the tcp on")

	flag.Parse()

	log.Info().Msg("Starting scap" + *webServerPtr + *hostPtr + *portPtr)

	if *modePtr {
		switch *webServerPtr {
		case "fiber":
			webserver.Fiber(*hostPtr, *portPtr)
			return
		case "echo":
			webserver.Echo(*hostPtr, *portPtr)
			return
		case "gin":
			webserver.Gin(*hostPtr, *portPtr)
			return
		case "iris":
			webserver.Iris(*hostPtr, *portPtr)
			return
		default:
			webserver.NetHttp(*hostPtr, *portPtr)
			return
		}
	} else {
		log.Info().Msg("Running in consumer mode")
		nsq.CreateConsumer()
	}
}
