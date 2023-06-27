package utils

import (
	"github.com/rs/zerolog/log"
	"time"
)

// TimeStamp is a struct to hold the execution time of a function
type TimeStamp struct {
	Instance string `json:"instance,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}

// TimeTrack is a utility function to measure the execution time of a function
func TimeTrack(start time.Time, instance string) TimeStamp {
	elapsed := time.Since(start)
	ts := TimeStamp{Duration: elapsed.Nanoseconds(), Instance: instance}
	log.Info().Interface("timestamp", ts).Msg("TimeTrack")
	return ts
}
