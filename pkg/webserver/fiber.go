package webserver

import (
	"encoding/json"
	"github.com/valerius21/scap/pkg/dto"
	"github.com/valerius21/scap/pkg/nsq"
	"github.com/valerius21/scap/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Fiber is a webserver implementation using the Fiber framework on top of fastHTTP
func Fiber(receiverHost, receiverPort string) {
	// Create a new Fiber instance
	app := fiber.New()

	// Define the routes and their corresponding handlers
	app.Get("/image", func(c *fiber.Ctx) error {
		defer utils.TimeTrack(time.Now(), "Fiber:ImageHandler")

		return c.SendString("imageHandler: not implemented")
	})
	app.Get("/empty", handler)
	//	func(c *fiber.Ctx) error {
	//	defer utils.TimeTrack(time.Now(), "Fiber:EmptyHandler")
	//
	//	// Send and receive messages here:
	//
	//	return nil
	//})
	//app.Get("/math", func(c *fiber.Ctx) error {
	//	defer utils.TimeTrack(time.Now(), "Fiber:MathHandler")
	//	s := makeSender()
	//	number := c.QueryInt("number")
	//	_, res, err := s.Send("math:" + strconv.Itoa(number))
	//	if err != nil {
	//		return err
	//	}
	//	return c.SendString(fmt.Sprintf("mathHandler: %s", res))
	//})
	//app.Get("/sleep", func(c *fiber.Ctx) error {
	//	defer utils.TimeTrack(time.Now(), "Fiber:SleepHandler")
	//	s := makeSender()
	//	_, _, err := s.Send("sleep:_")
	//	if err != nil {
	//		return err
	//	}
	//	return c.SendString("sleepHandler: slept for 1 second")
	//})

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
		return
	}
}
func handler(c *fiber.Ctx) error {
	startFunction := time.Now()
	// Create a message struct
	message := dto.Message{
		Name:     "empty",
		Args:     "",
		Duration: -1,
	}

	// Marshal the message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Error().Err(err).Msg("Error when marshaling the message")
		return err
	}

	startTrip := time.Now()
	// Publish the message to NSQ
	err = nsq.PublishMessage(messageBytes)
	if err != nil {
		log.Error().Err(err).Msg("Error when publishing the message to NSQ")
		return err
	}

	// Wait for the response from NSQ
	response, err := nsq.WaitForResponse()
	if err != nil {
		log.Error().Err(err).Msg("Error when waiting for the response from NSQ")
		return err
	}
	endTripTs := utils.TimeTrack(startTrip, "Fiber:EmptyHandler:Trip")

	// Unmarshal the response
	var resp dto.Message
	err = json.Unmarshal([]byte(response), &resp)

	ts := utils.TimeTrack(startFunction, "Fiber:EmptyHandler")
	wsResp := dto.WebServerResponse{
		Name:       "fiber:handler:empty",
		Args:       "",
		Message:    resp,
		TimeStamps: []utils.TimeStamp{ts, endTripTs},
	}
	wsRespBytes, err := json.Marshal(wsResp)
	if err != nil {
		log.Error().Err(err).Msg("Error when marshaling the response")
		return err
	}
	return c.Send(wsRespBytes)
}
