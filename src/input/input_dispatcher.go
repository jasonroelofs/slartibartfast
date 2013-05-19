package input

import (
	"events"
	"fmt"
	"github.com/go-gl/glfw"
)

type eventCallbackMap map[events.EventType]func(events.Event)
type keyEventMap map[int]events.EventType

type InputDispatcher struct {
	callbacks   eventCallbackMap
	keyMappings keyEventMap

	// List of events received. Gets cleared when requested.
	storedEvents EventList
}

func NewInputDispatcher() *InputDispatcher {
	mapper := InputDispatcher{
		callbacks:   make(eventCallbackMap),
		keyMappings: make(keyEventMap),
	}

	// Set up testing key mappings
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyEsc, events.Quit)

	glfw.SetKeyCallback(mapper.keyCallback)

	return &mapper
}

// On registers a callback to be called in the occurance of an event of type EventType.
// The callback will include event details, including key hit, and whether the key was
// pressed or released
// Use this method when you want input events outside of an Entity's Input component
func (self *InputDispatcher) On(event events.EventType, callback func(events.Event)) {
	self.callbacks[event] = callback
}

func (self *InputDispatcher) mapKeyToEvent(key int, eventType events.EventType) {
	self.keyMappings[key] = eventType
}

func (self *InputDispatcher) keyCallback(key, state int) {
	fmt.Println("Key pressed! ", key, state, string(key))

	if eventFromKey, ok := self.keyMappings[key]; ok {
		event := events.Event{
			EventType: eventFromKey,
			Pressed:   state == 1,
		}

		self.storedEvents = append(self.storedEvents, event)
		self.fireLocalCallback(event)
	}
}

func (self *InputDispatcher) fireLocalCallback(event events.Event) {
	eventToCallback := self.callbacks[event.EventType]
	if eventToCallback != nil {
		eventToCallback(event)
	}
}

// RecentEvents returns the list of events queued up since the last time
// this method was called. This method returns a copy of the events list
// then clears out it's internal list for the next frame.
func (self *InputDispatcher) RecentEvents() EventList {
	eventsList := self.storedEvents
	self.storedEvents = EventList{}
	return eventsList
}
