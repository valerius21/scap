package main

import (
	"flag"
	"github.com/valerius21/scap/pkg/receiver"
	"github.com/valerius21/scap/pkg/webserver"
)

func main() {
	webServerPtr := flag.String("webserver", "fiber", "the webserver to run")
	modePtr := flag.Bool("http", false, "if true, run the webserver in http mode,"+
		" otherwise operator mode")
	hostPtr := flag.String("host", "localhost", "the host to run the tcp on")
	portPtr := flag.String("port", "3333", "the port to run the tcp on")

	flag.Parse()

	if *modePtr {
		switch *webServerPtr {
		case "fiber":
			webserver.Fiber()
			return
		case "echo":
			webserver.Echo()
			return
		case "gin":
			webserver.Gin()
			return
		case "iris":
			webserver.Iris()
			return
		default:
			webserver.NetHttp()
			return
		}
	} else {
		receiver.CreateConnection(*hostPtr, *portPtr)
	}
}
