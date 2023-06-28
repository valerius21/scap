package nsq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/valerius21/scap/pkg/dto"
	"log"
	"time"
)

func CreateProducer(name, content, topic string) {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig() //Creating the Producer using NSQD Address
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	} //Init topic name and message
	msg := dto.Message{
		Name:      name,
		Args:      content,
		Timestamp: time.Now().String(),
	} //Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	} //Publish the Message
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}
