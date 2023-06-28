package dto

type Message struct {
	Name     string `json:"name"`
	Args     string `json:"args"`
	Duration int64  `json:"duration"`
}
