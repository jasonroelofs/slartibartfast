package util

import (
	"fmt"
	"log"
	"time"
)

// StartTiming returns the current time and a message built from the params sent through Sprintf.
// This should be paired with EndTiming in a defer() to set up method timing pair:
//
//   func SomeMethod() {
//     defer util.EndTiming(util.StartTiming("SomeMethod", ...))
//   }
//
func StartTiming(format string, parts ...interface{}) (startTime time.Time, message string) {
	message = fmt.Sprintf(format, parts...)
	startTime = time.Now()
	return
}

// EndTiming Prints the time since the given startTime, including the message.
func EndTiming(startTime time.Time, message string) {
	log.Printf("%s took %s\n", message, time.Since(startTime))
}
