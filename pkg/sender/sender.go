package sender

import "net"

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

// Send sends a message to the receiver
func (s *Sender) Send(msg string) (int, error) {
	return (*s.C).Write([]byte(msg))
}
