package webserver

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/dto"
	"github.com/valerius21/scap/pkg/nsq"
	"github.com/valerius21/scap/pkg/utils"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func CreateHandler(framework, handler, args string) ([]byte, error) {
	startFunction := time.Now()
	// Create a message struct
	message := dto.Message{
		Name:     handler,
		Data:     args,
		Duration: -1,
	}

	// Marshal the message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Error().Err(err).Msg("Error when marshaling the message")
		return nil, err
	}

	startTrip := time.Now()
	// Publish the message to NSQ
	err = nsq.PublishMessage(messageBytes)
	if err != nil {
		log.Error().Err(err).Msg("Error when publishing the message to NSQ")
		return nil, err
	}

	// Wait for the response from NSQ
	response, err := nsq.WaitForResponse()
	if err != nil {
		log.Error().Err(err).Msg("Error when waiting for the response from NSQ")
		return nil, err
	}
	endTripTs := utils.TimeTrack(startTrip, framework+":"+handler+":trip")

	// Unmarshal the response
	var resp dto.Message
	err = json.Unmarshal([]byte(response), &resp)

	ts := utils.TimeTrack(startFunction, framework+":"+handler+":function")
	wsResp := dto.WebServerResponse{
		Name:       framework + ":handler:" + handler,
		Args:       "",
		Message:    resp,
		TimeStamps: []utils.TimeStamp{ts, endTripTs},
	}
	wsRespBytes, err := json.Marshal(wsResp)
	if err != nil {
		log.Error().Err(err).Msg("Error when marshaling the response")
		return nil, err
	}
	return wsRespBytes, nil
}

// ImageSaver saves the image to a temporary file
func ImageSaver(file *multipart.FileHeader) (string, error) {
	fileName := "/tmp/scap/" + file.Filename

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileName, nil
}
