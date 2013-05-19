package input

import (
	"events"
)

type EventList []events.Event

type InputQueue interface {
	// RecentEvents returns the list of events that have fired since the
	// last time this method was called.
	RecentEvents() EventList
}
