package dto

import "github.com/valerius21/scap/pkg/utils"

type WebServerResponse struct {
	Name       string            `json:"name"`
	Args       string            `json:"args"`
	Message    Message           `json:"message"`
	TimeStamps []utils.TimeStamp `json:"time_stamps"`
}
