package input

import (
	"events"
	"log"
)

type eventCallback func(events.Event)
type eventCallbackList []eventCallback

type eventCallbackMap map[events.EventType]eventCallbackList
type keyCallbackMap map[int]eventCallbackList

type keyEventMap map[int][]events.EventType
type eventKeyMap map[events.EventType][]int
type eventTypeSet map[events.EventType]bool

// Testing tool, will remove once I have all key -> event mappings
// coming in via a file instead of hard-coded.
var InputDispatcherTesting bool

// InputDispatcher hooks into an InputEmitter and dispatches keyboard, mouse, and joystick events.
// It works via callbacks mostly but supports polling as well.
// Keys are mapped to Events, and Events are then used throughout the system.
//
// Mouse movement callback receives X and Y, which are pixel distances from the center of the screen.
// Positive right and up, Negative left and down.
//
// TODO This struct feels very heavy, look into ways of splitting up some of the
// responsibilities.
type InputDispatcher struct {
	emitter InputEmitter

	// User-specified callbacks according to an Event
	// or a specific raw Key event
	callbacks    eventCallbackMap
	keyCallbacks keyCallbackMap

	// Double mapping of how Key -> Event and Event -> Key
	// for easy look-up of either direction
	keyToEventMappings keyEventMap
	eventToKeyMappings eventKeyMap

	// List of events received. Gets cleared when requested.
	storedEvents EventList

	// Set of events that need to be polled
	pollingEvents eventTypeSet
}

func NewInputDispatcher(emitter InputEmitter) *InputDispatcher {
	mapper := InputDispatcher{
		emitter:            emitter,
		callbacks:          make(eventCallbackMap),
		keyCallbacks:       make(keyCallbackMap),
		keyToEventMappings: make(keyEventMap),
		eventToKeyMappings: make(eventKeyMap),
		pollingEvents:      make(eventTypeSet),
	}

	if !InputDispatcherTesting {
		// Set up testing key mappings
		mapper.mapKeyToEvent(KeyQ, events.Quit)
		mapper.mapKeyToEvent(KeyEsc, events.Quit)

		// MY FPS Movement mapping. Screw this WASD crap :P
		// Will move this to be something read from the settings file
		mapper.mapKeyToEvent(KeyE, events.MoveForward)
		mapper.mapKeyToEvent(KeyD, events.MoveBackward)
		mapper.mapKeyToEvent(KeyS, events.MoveLeft)
		mapper.mapKeyToEvent(KeyF, events.MoveRight)
		mapper.mapKeyToEvent(KeyW, events.TurnLeft)
		mapper.mapKeyToEvent(KeyR, events.TurnRight)

		mapper.mapKeyToEvent(KeyE, events.PanUp)
		mapper.mapKeyToEvent(KeyD, events.PanDown)
		mapper.mapKeyToEvent(KeyS, events.PanLeft)
		mapper.mapKeyToEvent(KeyF, events.PanRight)

		mapper.mapKeyToEvent(KeyI, events.ZoomIn)
		mapper.mapKeyToEvent(KeyO, events.ZoomOut)
	}

	emitter.KeyCallback(mapper.keyCallback)

	emitter.MousePositionCallback(mapper.mouseMoveCallback)
	emitter.MouseWheelCallback(mapper.mouseWheelCallback)
	emitter.MouseButtonCallback(mapper.mouseButtonCallback)

	return &mapper
}

// On registers a callback to be called in the occurance of an event of type EventType.
// The callback will include event details, including key hit, and whether the key was
// pressed or released
// Use this method when you want input events outside of an Entity's Input component
func (self *InputDispatcher) On(event events.EventType, callback func(events.Event)) {
	self.callbacks[event] = append(self.callbacks[event], callback)
}

// OnKey registers a callback for when a raw key event happens.
// Use this when you don't want to deal with the events mapping system and just want
// to watch for a key press. Should not be used with anything players will use.
func (self *InputDispatcher) OnKey(key int, callback func(events.Event)) {
	self.keyCallbacks[key] = append(self.keyCallbacks[key], callback)
}

// RecentEvents returns the list of events queued up since the last time
// this method was called. This method returns a copy of the events list
// then clears out it's internal list for the next frame.
// :: InputQueue
func (self *InputDispatcher) RecentEvents() EventList {
	eventsList := self.storedEvents
	polledEvents := self.findPolledEvents()
	eventsList = append(eventsList, polledEvents...)

	self.storedEvents = EventList{}
	return eventsList
}

// PollEvents adds the given events to the list of events this dispatcher
// should be polling for every frame
// :: InputQueue
func (self *InputDispatcher) PollEvents(eventsAdding []events.EventType) {
	for _, eventType := range eventsAdding {
		self.pollingEvents[eventType] = true
	}
}

// UnpollEvents removes the given events to the list of events this dispatcher
// should be polling for every frame
// :: InputQueue
func (self *InputDispatcher) UnpollEvents(eventsRemoving []events.EventType) {
	for _, toRemove := range eventsRemoving {
		delete(self.pollingEvents, toRemove)
	}
}

func (self *InputDispatcher) findPolledEvents() (polledEvents []events.Event) {
	var eventType events.EventType
	var eventKeys []int
	var eventKey int
	var ok bool

	for eventType, _ = range self.pollingEvents {
		eventKeys, ok = self.eventToKeyMappings[eventType]
		if !ok {
			continue
		}

		for _, eventKey = range eventKeys {
			if self.emitter.IsKeyPressed(eventKey) {
				polledEvents = append(polledEvents, events.Event{
					EventType: eventType,
					Pressed:   true,
				})
			}

			break
		}
	}

	return
}

func (self *InputDispatcher) mapKeyToEvent(key int, eventType events.EventType) {
	self.eventToKeyMappings[eventType] = append(self.eventToKeyMappings[eventType], key)
	self.keyToEventMappings[key] = append(self.keyToEventMappings[key], eventType)
}

func (self *InputDispatcher) keyCallback(key int, state KeyState) {
	log.Println("Key pressed! ", key, state, string(key))

	self.processKeyCallback(key, state)
	self.processEventCallbacks(key, state)
}

func (self *InputDispatcher) processKeyCallback(key int, state KeyState) {
	if callbacksFromKey, ok := self.keyCallbacks[key]; ok {
		for _, callback := range callbacksFromKey {
			callback(events.Event{Pressed: state == 1})
		}
	}
}

func (self *InputDispatcher) processEventCallbacks(key int, state KeyState) {
	if eventsFromKey, ok := self.keyToEventMappings[key]; ok {
		for _, eventFromKey := range eventsFromKey {
			event := events.Event{
				Pressed:   state == 1,
				EventType: eventFromKey,
			}

			self.storedEvents = append(self.storedEvents, event)
			self.fireLocalCallback(event)
		}
	}
}

func (self *InputDispatcher) fireLocalCallback(event events.Event) {
	eventsToCallback := self.callbacks[event.EventType]
	for _, eventCallback := range eventsToCallback {
		eventCallback(event)
	}
}

func (self *InputDispatcher) mouseMoveCallback(x, y int) {
	event := events.Event{
		EventType:  events.MouseMove,
		MouseXDiff: x,
		MouseYDiff: y,
	}

	log.Println("Mouse Moved", event)

	self.storedEvents = append(self.storedEvents, event)
}

func (self *InputDispatcher) mouseButtonCallback(button int, state KeyState) {
	log.Println("Mouse Button pressed", button, state)
}

func (self *InputDispatcher) mouseWheelCallback(position int) {
	log.Println("Mouse wheel", position)
}
