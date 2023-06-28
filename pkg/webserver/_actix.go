package webserver

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/sender"
	"net"
)

// Actix Webserver interfacing with the fastest Rust-Based Framework
func Actix(receiverHost, receiverPort string) error {

	// Create a connection to the receiver
	s := sender.CreateSender(receiverHost, receiverPort)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing connection")
		}
	}(*s.C)

	return fmt.Errorf("Not implemented")
}
