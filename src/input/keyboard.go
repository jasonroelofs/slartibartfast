package input

import (
	"events"
	"log"
)

type keyEventMap map[KeyCode][]events.EventType
type eventKeyMap map[events.EventType][]KeyCode

type keyCallbackMap map[KeyCode]eventCallbackList

// Keyboard handles mapping of Key inputs to Events
type Keyboard struct {
	// implements input.Input

	emitter InputEmitter

	// Mapping of keys to Events and vis versa
	keyToEvent keyEventMap
	eventToKey eventKeyMap

	// Callbacks hard-linked to specific keys
	keyCallbacks keyCallbackMap

	// Called when we receive an input that maps to an event
	eventReceived func(events.Event)
}

func NewKeyboard(emitter InputEmitter) *Keyboard {
	input := Keyboard{
		emitter:      emitter,
		keyToEvent:   make(keyEventMap),
		eventToKey:   make(eventKeyMap),
		keyCallbacks: make(keyCallbackMap),
	}

	emitter.KeyCallback(input.keyCallback)

	return &input
}

// Given a string of the key to map, hook it up to the event type given.
// If the given string does not map to a specific key, this method does nothing.
func (self *Keyboard) Map(inputName string, eventType events.EventType) {
	keyCode := KeyFromName(inputName)
	if keyCode != KeyNone {
		self.eventToKey[eventType] = append(self.eventToKey[eventType], keyCode)
		self.keyToEvent[keyCode] = append(self.keyToEvent[keyCode], eventType)
	}
}

// Install callback hook to be called when input is received that maps to an Event
func (self *Keyboard) OnEvent(callback func(events.Event)) {
	self.eventReceived = callback
}

// PollEvents looks at the current state of the keyboard, finding any keys who
// are mapped to the given list of events and if an Event should be built according
// to the state of the keys.
func (self *Keyboard) PollEvents(eventsToPoll EventTypeList) EventList {
	var polledEvents EventList
	var eventType events.EventType
	var eventKeys []KeyCode
	var eventKey KeyCode
	var ok bool

	for _, eventType = range eventsToPoll {
		eventKeys, ok = self.eventToKey[eventType]
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

	return polledEvents
}

// OnKey registers a callback for when a raw key event happens.
// Use this when you don't want to deal with the events mapping system and just want
// to watch for a key press.
func (self *Keyboard) OnKey(key KeyCode, callback func(events.Event)) {
	self.keyCallbacks[key] = append(self.keyCallbacks[key], callback)
}

func (self *Keyboard) keyCallback(key KeyCode, state KeyState) {
	log.Println("Key pressed! ", key, state, string(key))

	event := events.Event{
		Pressed:  state == KeyPressed || state == KeyRepeated,
		Repeated: state == KeyRepeated,
	}

	self.fireKeyCallbacks(key, event)
	self.registerNewEvents(key, event)
}

func (self *Keyboard) fireKeyCallbacks(key KeyCode, event events.Event) {
	if callbacksFromKey, ok := self.keyCallbacks[key]; ok {
		for _, callback := range callbacksFromKey {
			callback(event)
		}
	}
}

func (self *Keyboard) registerNewEvents(key KeyCode, event events.Event) {
	eventsFromKey, eventFound := self.keyToEvent[key]
	if eventFound {
		for _, eventFromKey := range eventsFromKey {
			event.EventType = eventFromKey
			self.eventReceived(event)
		}
	}
}
