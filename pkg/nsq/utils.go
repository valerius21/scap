package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/valerius21/scap/pkg/common"
)

// WaitForResponse waits for a response message from NSQ
func WaitForResponse() (string, error) {
	// Instantiate a new NSQ consumer for response messages
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(common.ResponseTopic, common.ResponseChannel, config)
	if err != nil {
		return "", fmt.Errorf("error creating NSQ consumer for response messages: %w", err)
	}

	responseReceived := make(chan string)

	// Set the Handler for response messages
	consumer.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		response := string(m.Body)
		responseReceived <- response
		return nil
	}))

	// Connect to NSQD for response messages
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		return "", fmt.Errorf("error connecting to NSQD for response messages: %w", err)
	}

	// Wait for the response message
	response := <-responseReceived

	// Stop the consumer
	consumer.Stop()

	return response, nil
}

// PublishMessage publishes a message to NSQ
func PublishMessage(message []byte) error {
	// Instantiate a new NSQ producer
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		return fmt.Errorf("error creating NSQ producer: %w", err)
	}

	// Publish the message to a topic
	err = producer.Publish(common.DefaultTopic, message)
	if err != nil {
		return fmt.Errorf("error publishing message to NSQ: %w", err)
	}

	// Gracefully stop the producer
	producer.Stop()

	return nil
}
