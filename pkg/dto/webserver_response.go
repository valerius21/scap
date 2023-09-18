package dto

import "github.com/valerius21/scap/pkg/utils"

// WebServerResponse is the struct that is sent to the client
type WebServerResponse struct {
	// Name is the name of the handler-function
	Name string `json:"name"`
	// Args are the arguments that were passed to the handler-function
	Args string `json:"args"`
	// Message is the message with all the trip and execution data that was sent to the client
	Message Message `json:"message"`
	// TimeStamps is an array of TimeStamp structs that contain the trip and execution-duration data
	TimeStamps []utils.TimeStamp `json:"time_stamps"`
}
