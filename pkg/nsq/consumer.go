package nsq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/common"
	"github.com/valerius21/scap/pkg/dto"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type messageHandler struct{}

func DefaultStopChannel() {
	// Gracefully handle SIGINT and SIGTERM signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func CreateConsumer(topic, channel string, stopper func()) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig() //Tweak several common setup in config
	// Maximum number of times this consumer will attempt to process a message before giving up
	config.MaxAttempts = 10 // Maximum number of messages to allow in flight
	config.MaxInFlight = 5  // Maximum duration when REQueueing
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0 //Init topic name and channel
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error when creating the consumer")
	} // Set the Handler for messages received by this Consumer.
	consumer.AddHandler(&messageHandler{}) //Use nsqlookupd to find nsqd instances

	// Connect to NSQD
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal().Err(err).Msg("Error when connecting to NSQD")
	}

	stopper()

	// Stop the consumer
	consumer.Stop()
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error { //Process the Message
	var request dto.Message
	if err := json.Unmarshal(m.Body, &request); err != nil {
		log.Error().Err(err).Msg("Error when Unmarshaling the message body")
		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return err
	} //Print the Message

	log.Info().Msg("Message")
	log.Info().Msg("--------------------")
	log.Info().Msgf("Name : %s", request.Name)
	log.Info().Msgf("Args : %s", request.Args)
	log.Info().Msgf("Timestamp : %s", request.Timestamp)
	log.Info().Msg("--------------------")
	log.Info().Msgf("%s", request) // Will automatically set the message as finish

	CreateProducer(request.Name, request.Args, common.ResponseTopic)

	return nil
}
