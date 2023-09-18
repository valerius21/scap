package dto

type Message struct {
	// Name is the function-name
	Name string `json:"name"`
	// Data are the arguments to the function
	Data string `json:"data"`
	// Duration is the time in nanoseconds that the function took to execute
	Duration int64 `json:"duration"`
}
