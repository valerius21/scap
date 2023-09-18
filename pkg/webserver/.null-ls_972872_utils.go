package webserver

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/rpc"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/valerius21/scap/pkg/dto"
	_ "github.com/valerius21/scap/pkg/rpc_services"
	"github.com/valerius21/scap/pkg/utils"
)

const SERIVCE_NAME = "rpc_services.HandlerService.HandleMessage"

// TODO: rename args to their actual name
func CreateHandler(framework, fn, fnArgs string) ([]byte, error) {
	startFunction := time.Now()

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		return nil, err
	}

	defer client.Close()

	var response dto.Message

	startTrip := time.Now()

	err = client.Call("HandlerService.HandleMessage", &dto.Message{
		Name:     fn,
		Data:     fnArgs,
		Duration: -1,
	}, &response)
	if err != nil {
		return nil, err
	}

	endTripTs := utils.TimeTrack(startTrip, framework+":"+fn+":trip")

	ts := utils.TimeTrack(startFunction, framework+":"+fn+":function")

	wsResp := dto.WebServerResponse{
		Name:       framework + ":handler:" + fn,
		Args:       "",
		Message:    response,
		TimeStamps: []utils.TimeStamp{ts, endTripTs},
	}
	wsRespBytes, err := json.Marshal(wsResp)
	if err != nil {
		log.Error().Err(err).Msg("Error when marshaling the response")
		return nil, err
	}

  // Return the captured data
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
