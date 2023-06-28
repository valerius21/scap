package receiver

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/fns"
	"github.com/valerius21/scap/pkg/utils"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// this package is responsible for receiving messages from the sender
// toggle between the different webserver implementations

// CreateConnection creates a tcp connection on the specified host and port
func CreateConnection(host, port string) net.Conn {
	l, err := net.Listen("tcp", host+":"+port)

	if err != nil {
		panic(err)
	}

	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			panic(err)
		}
	}(l)
	log.Info().Msg("Listening on " + host + ":" + port)
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
		n := fns.MathFn(number)
		utils.TimeTrack(now, "TCP:MathHandler")
		conn.Write([]byte(fmt.Sprintf("Result: %.2f", n)))
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
		conn.Write([]byte("Done"))
		utils.TimeTrack(now, "TCP:EmptyHandler")
	} else {
		// Handle unknown function
		log.Info().Msg("Unknown function")
		conn.Write([]byte("Unknown function"))
	}
}
