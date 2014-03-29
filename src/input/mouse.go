package input

import (
	"events"
	"log"
)

type mouseButtonEventMap map[MouseButtonCode][]events.EventType
type eventMouseButtonMap map[events.EventType][]MouseButtonCode

// Mouse handles mapping of Button inputs to Events
type Mouse struct {
	// implements input.Input

	emitter InputEmitter

	// Mapping of buttons to Events and vis versa
	buttonToEvent mouseButtonEventMap
	eventToButton eventMouseButtonMap

	// Called when we receive an input that maps to an event
	eventReceived func(events.Event)
}

func NewMouse(emitter InputEmitter) *Mouse {
	input := Mouse{
		emitter:         emitter,
		buttonToEvent:   make(mouseButtonEventMap),
		eventToButton:   make(eventMouseButtonMap),
	}

	emitter.MousePositionCallback(input.mouseMoveCallback)
	emitter.MouseWheelCallback(input.mouseWheelCallback)
	emitter.MouseButtonCallback(input.mouseButtonCallback)

	return &input
}

// Given a string of the button to map, hook it up to the event type given.
// If the given string does not map to a specific button, this method does nothing.
func (self *Mouse) Map(inputName string, eventType events.EventType) {
	buttonCode := MouseButtonFromName(inputName)
	if buttonCode != MouseNone {
		self.eventToButton[eventType] = append(self.eventToButton[eventType], buttonCode)
		self.buttonToEvent[buttonCode] = append(self.buttonToEvent[buttonCode], eventType)
	}
}

// Install callback hook to be called when input is received that maps to an Event
func (self *Mouse) OnEvent(callback func(events.Event)) {
	self.eventReceived = callback
}

// No-op for Mouse right now
func (self *Mouse) PollEvents(eventsToPoll EventTypeList) EventList {
	return EventList{}
}


func (self *Mouse) mouseMoveCallback(x, y int) {
	event := events.Event{
		EventType:  events.MouseMove,
		MouseXDiff: x,
		MouseYDiff: y,
	}

	log.Println("Mouse Moved", event)

	self.eventReceived(event)
}

func (self *Mouse) mouseButtonCallback(mouseButton MouseButtonCode, state KeyState) {
	log.Println("Mouse Button pressed", mouseButton, state)

	event := events.Event{
		Pressed:  state == KeyPressed || state == KeyRepeated,
		Repeated: state == KeyRepeated,
	}

	self.registerNewEvents(mouseButton, event)
}

func (self *Mouse) registerNewEvents(button MouseButtonCode, event events.Event) {
	eventsFromButton, eventFound := self.buttonToEvent[button]
	if eventFound {
		for _, eventFromButton := range eventsFromButton {
			event.EventType = eventFromButton
			self.eventReceived(event)
		}
	}
}

func (self *Mouse) mouseWheelCallback(position int) {
	log.Println("Mouse wheel", position)
}
