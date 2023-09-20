package main

import (
	"flag"
	"net"
	"net/rpc"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/config"
	"github.com/valerius21/scap/pkg/rpc_services"
	"github.com/valerius21/scap/pkg/webserver"
)

func main() {
	runLogFile, _ := os.OpenFile(
		"scap.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	if _, err := os.Stat("/tmp/scap"); os.IsNotExist(err) {
		err := os.Mkdir("/tmp/scap", 0o644) // Owner -> RW, Others -> R
		if err != nil {
			log.Error().Err(err).Msg("Error when creating the tmp directory")
			return
		}
	}

	configPathPtr := flag.String("config", "/tmp/scap/config.yaml", "the path to the config file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPathPtr)
	if err != nil {
		panic(err)
	}

	if cfg.Server.IsServer {
		log.Info().Msg("Running in server mode")
		handlerService := new(rpc_services.HandlerService)
		rpc.Register(handlerService)

		listener, err := net.Listen("tcp", cfg.EmitterAddress)
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

		switch cfg.Client.WebServer {
		case "fiber":
			webserver.Fiber(cfg.Client.Port)
			return
		case "echo":
			webserver.Echo(cfg.Client.Port)
			return
		case "gin":
			webserver.Gin(cfg.Client.Port)
			return
		case "iris":
			webserver.Iris(cfg.Client.Port)
			return
		default:
			webserver.NetHttp(cfg.Client.Port)
			return
		}

	}
}
