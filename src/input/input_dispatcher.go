package input

import (
	"configs"
	"events"
)

type eventCallback func(events.Event)
type eventCallbackList []eventCallback

type eventCallbackMap map[events.EventType]eventCallbackList
type eventTypeSet map[events.EventType]bool

// InputDispatcher dispatches Input events based on keyboard, mouse, and joystick input.
//
// This object supports callbacks and per-frame polling and key/button specific callbacks when
// you don't want to work with the Events system.
//
// Movement events are pixel distances from the center of the screen.
// Positive right and up, Negative left and down.
type InputDispatcher struct {
	// Implements input.InputQueue

	config *configs.Config

	// List of Inputs we're receiving events from
	inputs []Input

	// User-specified callbacks according to an Event
	// or a specific raw Key event
	callbacks eventCallbackMap

	// List of events received. Gets cleared when requested.
	recentEvents EventList

	// Set of events that need to be polled
	pollingEvents eventTypeSet
}

func NewInputDispatcher(config *configs.Config, inputs ...Input) *InputDispatcher {
	mapper := InputDispatcher{
		config:        config,
		inputs:        inputs,
		callbacks:     make(eventCallbackMap),
		pollingEvents: make(eventTypeSet),
	}

	for _, input := range inputs {
		input.OnEvent(mapper.receiveEvent)
	}

	mapper.initializeBindings()

	return &mapper
}

func (self *InputDispatcher) initializeBindings() {
	inputBindings := make(map[string][]string)
	self.config.Get("input", &inputBindings)

	for eventName, inputList := range inputBindings {
		for _, inputName := range inputList {
			for _, handler := range self.inputs {
				handler.Map(inputName, events.EventFromName(eventName))
			}
		}
	}
}

// On registers a callback to be called in the occurance of an event of type EventType.
// The callback will include event details, including key hit, and whether the key was
// pressed or released
// Use this method when you want input events outside of an Entity's Input component
func (self *InputDispatcher) On(event events.EventType, callback func(events.Event)) {
	self.callbacks[event] = append(self.callbacks[event], callback)
}

// Called by Inputs when a new event comes in from the system
func (self *InputDispatcher) receiveEvent(event events.Event) {
	self.recentEvents = append(self.recentEvents, event)
	self.fireEventCallbacks(event)
}

func (self *InputDispatcher) fireEventCallbacks(event events.Event) {
	for _, eventCallback := range self.callbacks[event.EventType] {
		eventCallback(event)
	}
}

// RecentEvents returns the list of events queued up since the last time
// this method was called. This method returns a copy of the events list
// then clears out it's internal list for the next frame.
// :: InputQueue
func (self *InputDispatcher) RecentEvents() EventList {
	eventsList := self.recentEvents
	var eventsToPoll EventTypeList

	for eventType, _ := range self.pollingEvents {
		eventsToPoll = append(eventsToPoll, eventType)
	}

	for _, input := range self.inputs {
		eventsList = append(
			eventsList,
			input.PollEvents(eventsToPoll)...,
		)
	}

	self.recentEvents = EventList{}

	return eventsList
}

// PollEvents adds the given events to the list of events this dispatcher
// should be polling for every frame
// :: InputQueue
func (self *InputDispatcher) PollEvents(eventsAdding EventTypeList) {
	for _, eventType := range eventsAdding {
		self.pollingEvents[eventType] = true
	}
}

// UnpollEvents removes the given events to the list of events this dispatcher
// should be polling for every frame
// :: InputQueue
func (self *InputDispatcher) UnpollEvents(eventsRemoving EventTypeList) {
	for _, toRemove := range eventsRemoving {
		delete(self.pollingEvents, toRemove)
	}
}
