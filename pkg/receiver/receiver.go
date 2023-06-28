package receiver

import (
	"fmt"
	"github.com/pawelgaczynski/gain"
	"github.com/pawelgaczynski/gain/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync/atomic"
)

const (
	port            = 8888
	numberOfClients = 200
)

var testData = []byte("echo")

type EventHandler struct {
	server gain.Server

	logger zerolog.Logger

	overallBytesSent atomic.Uint64
}

func (e *EventHandler) OnStart(server gain.Server) {
	e.server = server
	e.logger = zerolog.New(os.Stdout).With().Logger().Level(zerolog.InfoLevel)
}

func (e *EventHandler) OnAccept(conn gain.Conn) {
	e.logger.Info().
		Int("active connections", e.server.ActiveConnections()).
		Str("remote address", conn.RemoteAddr().String()).
		Msg("New connection accepted")
}

func (e *EventHandler) OnRead(conn gain.Conn, n int) {
	e.logger.Info().
		Int("bytes", n).
		Str("remote address", conn.RemoteAddr().String()).
		Msg("Bytes received from remote peer")

	var (
		err    error
		buffer []byte
	)

	buffer, err = conn.Next(n)
	if err != nil {
		return
	}

	_, _ = conn.Write(buffer)
}

func (e *EventHandler) OnWrite(conn gain.Conn, n int) {
	buf := make([]byte, n)
	e.overallBytesSent.Add(uint64(n))

	e.logger.Info().
		Int("bytes", n).
		Str("remote address", conn.RemoteAddr().String()).
		Msg("Bytes sent to remote peer")

	n, err := conn.Write(buf)

	if err != nil {
		e.logger.Error().Err(err).Msg("Error during read")
	}

	log.Info().Msg(fmt.Sprintf("Message: %s", string(buf[:n])))

	err = conn.Close()
	if err != nil {
		e.logger.Error().Err(err).Msg("Error during connection close")
	}
}

func (e *EventHandler) OnClose(conn gain.Conn, err error) {
	log := e.logger.Info().
		Str("remote address", conn.RemoteAddr().String())
	if err != nil {
		log.Err(err).Msg("Connection from remote peer closed")
	} else {
		log.Msg("Connection from remote peer closed by server")
	}

	if e.overallBytesSent.Load() >= uint64(len(testData)*numberOfClients) {
		e.server.AsyncShutdown()
	}
}

func CreateReceiver() {
	err := gain.ListenAndServe(
		fmt.Sprintf("tcp://127.0.0.1:%d", port), &EventHandler{}, gain.WithLoggerLevel(logger.WarnLevel))
	if err != nil {
		log.Error().Err(err).Msg("Error during server startup")
	}
}
