package input

import (
	"events"
)

type EventList []events.Event

type InputQueue interface {
	// RecentEvents returns the list of events that have fired since the
	// last time this method was called.
	RecentEvents() EventList

	// PollEvents takes a list of EventTypes that should be polled against
	// (check the state of the related key(s) every frame).
	PollEvents([]events.EventType)

	// UnpollEvents undoes what PollEvents does, removes the given EventTypes
	// from the list of Events this queue is polling
	UnpollEvents([]events.EventType)
}
