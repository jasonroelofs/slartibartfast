package input

import (
	"events"
)

// Input types handle all mappings from the OS layer into
// the appropriate Events. Each Input should handle one aspect of the input
// system (keyboard, mouse, joystick, etc).
type Input interface {
	// Map a given Input code to a system EventType
	Map(inputName string, event events.EventType)

	// OnEvent sets a callback that will be fired whenever an Input event is
	// received from the underlying system.
	OnEvent(func(events.Event))

	// PollEvents takes a list of EventTypes to poll for and should return
	// a list of Events according to the current state of all inputs mapped
	// to the given EventTypes
	PollEvents(events.EventTypeList) events.EventList
}
