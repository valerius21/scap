package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/common"
	"github.com/valerius21/scap/pkg/nsq"
	"github.com/valerius21/scap/pkg/webserver"
)

func main() {
	webServerPtr := flag.String("webserver", "fiber", "the webserver to run")
	modePtr := flag.Bool("http", false, "if true, run the webserver in http mode,"+
		" otherwise operator mode")
	hostPtr := flag.String("nsqHost", "127.0.0.1", "the host of the nsq server")
	portPtr := flag.String("port", "3000", "the port to run the http on")

	flag.Parse()

	if *modePtr {
		log.Info().Msg("Starting SCAP (HTTP)")
	} else {
		log.Info().Msg("Starting SCAP (Node)")
	}

	common.NsqHost = *hostPtr

	if *modePtr {
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
	} else {
		log.Info().Msg("Running in consumer mode")
		nsq.CreateConsumer()
	}
}
