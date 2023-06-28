package nsq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/dto"
	"time"
)

func CreateProducer(name, content, topic string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig() //Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Error().Err(err).Msg("Error when creating the producer")
	} //Init topic name and message
	msg := dto.Message{
		Name:      name,
		Args:      content,
		Timestamp: time.Now().String(),
	} //Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Error().Err(err).Msg("Error when marshaling the message")
	} //Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Error().Err(err).Msg("Error when publishing the message")
	}
}
