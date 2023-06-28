package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/common"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type messageHandler struct {
	producer *nsq.Producer
}

func DefaultStopChannel() {
	// Gracefully handle SIGINT and SIGTERM signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func CreateConsumer() {
	// The only valid way to instantiate the Config
	config := nsq.NewConfig()
	// Tweak several common setup in config
	// Maximum number of times this consumer will attempt to process a message before giving up
	config.MaxAttempts = 10
	// Maximum number of messages to allow in flight
	config.MaxInFlight = 5
	// Maximum duration when REQueueing
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	// Init topic name and channel
	consumer, err := nsq.NewConsumer(common.DefaultTopic, common.DefaultChannel, config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error when creating the consumer")
	}

	// Instantiate a new NSQ producer
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal().Err(err).Msg("Error when creating the producer")
	}

	// Set the Handler for messages received by this Consumer.
	consumer.AddHandler(&messageHandler{
		producer: producer,
	})

	// Use nsqlookupd to find nsqd instances
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal().Err(err).Msg("Error when connecting to NSQD")
	}

	DefaultStopChannel()

	// Stop the consumer
	consumer.Stop()

	// Stop the producer
	producer.Stop()
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	// Process the Message
	log.Info().Msgf("Received message: %s", m.Body)

	n := time.Now().Second()
	i := strconv.Itoa(n)

	// Create a response message
	response := []byte("Response to the incoming message: " + i)

	// Publish the response message to a different topic
	err := h.producer.Publish(common.ResponseTopic, response)
	if err != nil {
		log.Error().Err(err).Msg("Failed to publish response message")
	}

	// Will automatically set the message as finished
	return nil
}
