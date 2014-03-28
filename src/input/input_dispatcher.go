package input

import (
	"configs"
	"events"
	"log"
)

type eventCallback func(events.Event)
type eventCallbackList []eventCallback

type eventCallbackMap map[events.EventType]eventCallbackList
type keyCallbackMap map[KeyCode]eventCallbackList

type keyEventMap map[KeyCode][]events.EventType
type eventKeyMap map[events.EventType][]KeyCode

type mouseButtonEventMap map[MouseButtonCode][]events.EventType
type eventMouseButtonMap map[events.EventType][]MouseButtonCode

type eventTypeSet map[events.EventType]bool

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
	config  *configs.Config
	emitter InputEmitter

	// User-specified callbacks according to an Event
	// or a specific raw Key event
	callbacks    eventCallbackMap
	keyCallbacks keyCallbackMap

	// Double mapping of how Key -> Event and Event -> Key
	// for easy look-up of either direction
	keyToEventMappings keyEventMap
	eventToKeyMappings eventKeyMap

	// Same for Mouse buttons
	mouseButtonToEventMappings mouseButtonEventMap
	eventToMouseButtonMappings eventMouseButtonMap

	// List of events received. Gets cleared when requested.
	storedEvents EventList

	// Set of events that need to be polled
	pollingEvents eventTypeSet
}

func NewInputDispatcher(config *configs.Config, emitter InputEmitter) *InputDispatcher {
	mapper := InputDispatcher{
		config:                     config,
		emitter:                    emitter,
		callbacks:                  make(eventCallbackMap),
		keyCallbacks:               make(keyCallbackMap),
		keyToEventMappings:         make(keyEventMap),
		eventToKeyMappings:         make(eventKeyMap),
		mouseButtonToEventMappings: make(mouseButtonEventMap),
		eventToMouseButtonMappings: make(eventMouseButtonMap),
		pollingEvents:              make(eventTypeSet),
	}

	mapper.initializeBindings()
	mapper.initializeCallbacks()

	return &mapper
}

func (self *InputDispatcher) initializeBindings() {
	inputBindings := make(map[string][]string)
	self.config.Get("input", &inputBindings)
	var key KeyCode
	var mouseButton MouseButtonCode

	for eventName, inputList := range inputBindings {
		for _, inputName := range inputList {
			key = KeyFromName(inputName)
			mouseButton = MouseButtonFromName(inputName)

			if key != KeyNone {
				self.mapKeyToEvent(key, events.EventFromName(eventName))
			}

			if mouseButton != MouseNone {
				self.mapMouseButtonToEvent(mouseButton, events.EventFromName(eventName))
			}
		}
	}
}

func (self *InputDispatcher) initializeCallbacks() {
	self.emitter.KeyCallback(self.keyCallback)

	self.emitter.MousePositionCallback(self.mouseMoveCallback)
	self.emitter.MouseWheelCallback(self.mouseWheelCallback)
	self.emitter.MouseButtonCallback(self.mouseButtonCallback)
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
func (self *InputDispatcher) OnKey(key KeyCode, callback func(events.Event)) {
	self.keyCallbacks[key] = append(self.keyCallbacks[key], callback)
}

// ShowCursor flags the cursor to be shown.
func (self *InputDispatcher) ShowCursor() {
	self.emitter.ShowCursor()
}

// HideCursor flags the cursor to be hidden and makes sure that the cursor's actual
// location stays constrained to the center of the screen.
func (self *InputDispatcher) HideCursor() {
	self.emitter.HideCursor()
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
	var eventKeys []KeyCode
	var eventKey KeyCode
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
					Repeated:  self.emitter.IsKeyRepeated(eventKey),
				})
			}

			break
		}
	}

	return
}

func (self *InputDispatcher) mapKeyToEvent(key KeyCode, eventType events.EventType) {
	self.eventToKeyMappings[eventType] = append(self.eventToKeyMappings[eventType], key)
	self.keyToEventMappings[key] = append(self.keyToEventMappings[key], eventType)
}

func (self *InputDispatcher) mapMouseButtonToEvent(button MouseButtonCode, eventType events.EventType) {
	self.eventToMouseButtonMappings[eventType] = append(self.eventToMouseButtonMappings[eventType], button)
	self.mouseButtonToEventMappings[button] = append(self.mouseButtonToEventMappings[button], eventType)
}

func (self *InputDispatcher) keyCallback(key KeyCode, state KeyState) {
	log.Println("Key pressed! ", key, state, string(key))
	event := events.Event{
		Pressed:  state == KeyPressed || state == KeyRepeated,
		Repeated: state == KeyRepeated,
	}

	self.processKeyCallback(key, event)
	self.processEventCallbacks(key, event)
}

func (self *InputDispatcher) processKeyCallback(key KeyCode, event events.Event) {
	if callbacksFromKey, ok := self.keyCallbacks[key]; ok {
		for _, callback := range callbacksFromKey {
			callback(event)
		}
	}
}

func (self *InputDispatcher) processEventCallbacks(key KeyCode, event events.Event) {
	if eventsFromKey, ok := self.keyToEventMappings[key]; ok {
		for _, eventFromKey := range eventsFromKey {
			event.EventType = eventFromKey

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

func (self *InputDispatcher) mouseButtonCallback(mouseButton MouseButtonCode, state KeyState) {
	log.Println("Mouse Button pressed", mouseButton, state)

	event := events.Event{
		Pressed:  state == KeyPressed || state == KeyRepeated,
		Repeated: state == KeyRepeated,
	}

	if eventsFromMouseButton, ok := self.mouseButtonToEventMappings[mouseButton]; ok {
		for _, eventFromMouseButton := range eventsFromMouseButton {
			event.EventType = eventFromMouseButton

			self.storedEvents = append(self.storedEvents, event)
			self.fireLocalCallback(event)
		}
	}
}

func (self *InputDispatcher) mouseWheelCallback(position int) {
	log.Println("Mouse wheel", position)
}
