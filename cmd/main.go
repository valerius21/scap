package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/fns"
	"github.com/valerius21/scap/pkg/utils"
	"github.com/valerius21/scap/pkg/webserver"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
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
		l, err := net.Listen("tcp", *hostPtr+":"+*portPtr)

		if err != nil {
			panic(err)
		}

		defer func(l net.Listener) {
			err := l.Close()
			if err != nil {
				panic(err)
			}
		}(l)
		log.Info().Msg("Listening on " + *hostPtr + ":" + *portPtr)
		for {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				log.Error().Err(err).Msg("Error accepting")
				os.Exit(1)
			}
			// Handle connections in a new goroutine.
			go handleRequest(conn)
		}
	}

}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Close the connection when you're done with it.
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing connection")
		}
	}(conn)

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	// Builds the message.
	n := bytes.Index(buf, []byte{0})
	message := string(buf[:n])

	// determine before ":" in message, if the substring belongs to "math", "sleep", "image" or "empty"
	// and call the corresponding function
	// Determine the substring before ":" in message
	substr := strings.Split(message, ":")[0]
	args := strings.Split(message, ":")[1]

	// Call the corresponding function based on the substring
	if substr == "math" {
		number, err := strconv.Atoi(args)
		if err != nil {
			log.Error().Err(err).Msg("Error converting string to int")
			return
		}
		now := time.Now()
		// Call math function
		res := fns.MathFn(number)
		ts := utils.TimeTrack(now, "TCP:MathHandler")

		gres := gin.H{
			"timestamp": ts,
			"result":    res,
		}
		gres.
		conn.Write([]byte(gres.))
	} else if substr == "sleep" {
		now := time.Now()
		// Call sleep function
		fns.SleeperFn(1)
		utils.TimeTrack(now, "TCP:SleepHandler")
		conn.Write([]byte("Done"))
	} else if substr == "image" {
		now := time.Now()
		// Call image function
		fns.GenerateImageMetadataFn()
		utils.TimeTrack(now, "TCP:ImageHandler")
	} else if substr == "empty" {
		now := time.Now()
		// Call empty function
		fns.EmptyFn()
		utils.TimeTrack(now, "TCP:EmptyHandler")
	} else {
		// Handle unknown function
		log.Info().Msg("Unknown function")
	}

	// Write the message in the connection channel.
	//conn.Write([]byte(message))
}
