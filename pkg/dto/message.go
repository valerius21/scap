package dto

type Message struct {
	Name     string `json:"name"`
	Data     string `json:"data"`
	Duration int64  `json:"duration"`
}
