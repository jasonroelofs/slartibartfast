package input

import (
	"events"
)

type InputQueue interface {
	// RecentEvents returns the list of events that have fired since the
	// last time this method was called.
	RecentEvents() events.EventList

	// PollEvents takes a list of EventTypes that should be polled against
	// (check the state of the related key(s) every frame).
	PollEvents(events.EventTypeList)

	// UnpollEvents undoes what PollEvents does, removes the given EventTypes
	// from the list of Events this queue is polling
	UnpollEvents(events.EventTypeList)
}
