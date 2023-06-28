package sender

import (
	"net"
	"time"
)

// Sender is an interface for sending messages to the webservers
type Sender struct {
	C *net.Conn
}

// CreateSender creates a tcp connection to the specified host and port
func CreateSender(host, port string) Sender {
	c, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	return Sender{
		C: &c,
	}
}

// Send sends a message to the receiver and waits for a response
func (s *Sender) Send(msg string) (int, string, error) {
	// Send the message
	bytesWritten, err := (*s.C).Write([]byte(msg))
	if err != nil {
		return 0, "", err
	}

	// Set a timeout for reading the response
	(*s.C).SetReadDeadline(time.Now().Add(5 * time.Second)) // Adjust the timeout duration as needed

	// Read the response
	buffer := make([]byte, 1024) // Adjust the buffer size as needed
	bytesRead, err := (*s.C).Read(buffer)
	if err != nil {
		return 0, "", err
	}

	// Process the response
	response := string(buffer[:bytesRead])

	return bytesWritten, response, nil
}
