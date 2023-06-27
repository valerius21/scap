package fns

import "time"

// SleeperFn is a function that sleeps for a given duration.
func SleeperFn(duration time.Duration) {
	time.Sleep(duration)
}
