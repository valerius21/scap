package fns

import "time"

// SleeperFn is a function that sleeps for 1 second.
func SleeperFn() {
	time.Sleep(1 * time.Second)
}
