package input

import (
	"events"
	"github.com/go-gl/glfw"
	"log"
)

type eventCallbackMap map[events.EventType]func(events.Event)
type keyCallbackMap map[int]func(events.Event)
type keyEventMap map[int]events.EventType

type InputDispatcher struct {
	callbacks    eventCallbackMap
	keyMappings  keyEventMap
	keyCallbacks keyCallbackMap

	// List of events received. Gets cleared when requested.
	storedEvents EventList

	// GLFW, when disabling the cursor, seems to end up triggering
	// a mouse-move event that is the distance mouse moved to be the
	// center of the window. This is making crazy swinging so I'm ignoring
	// it until I find a better way to handle this.
	firstMouseMoveIgnored bool
}

func NewInputDispatcher() *InputDispatcher {
	mapper := InputDispatcher{
		callbacks:    make(eventCallbackMap),
		keyMappings:  make(keyEventMap),
		keyCallbacks: make(keyCallbackMap),
	}

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

	glfw.SetKeyCallback(mapper.keyCallback)

	glfw.SetMouseButtonCallback(mapper.keyCallback)
	glfw.SetMousePosCallback(mapper.mouseCallback)
	glfw.SetMouseWheelCallback(mapper.mouseWheelCallback)

	glfw.Disable(glfw.MouseCursor)

	mapper.resetMouse()

	return &mapper
}

// On registers a callback to be called in the occurance of an event of type EventType.
// The callback will include event details, including key hit, and whether the key was
// pressed or released
// Use this method when you want input events outside of an Entity's Input component
func (self *InputDispatcher) On(event events.EventType, callback func(events.Event)) {
	self.callbacks[event] = callback
}

// OnKey registers a callback for when a raw key event happens. Use when there isn't an appropriate
// event defined or just for testing things out
func (self *InputDispatcher) OnKey(key int, callback func(events.Event)) {
	self.keyCallbacks[key] = callback
}

func (self *InputDispatcher) mapKeyToEvent(key int, eventType events.EventType) {
	self.keyMappings[key] = eventType
}

func (self *InputDispatcher) keyCallback(key, state int) {
	log.Println("Key pressed! ", key, state, string(key))

	event := events.Event{
		Pressed: state == 1,
	}

	if callbackFromKey, ok := self.keyCallbacks[key]; ok {
		callbackFromKey(event)
	}

	if eventFromKey, ok := self.keyMappings[key]; ok {
		event.EventType = eventFromKey

		self.storedEvents = append(self.storedEvents, event)
		self.fireLocalCallback(event)
	}
}

func (self *InputDispatcher) mouseCallback(x, y int) {
	if !self.firstMouseMoveIgnored {
		self.firstMouseMoveIgnored = true
		return
	}

	event := events.Event{
		EventType:  events.MouseMove,
		MouseXDiff: x,
		MouseYDiff: y,
	}

	log.Println("Mouse Moved", event)

	self.storedEvents = append(self.storedEvents, event)
	self.resetMouse()
}

func (self *InputDispatcher) resetMouse() {
	glfw.SetMousePos(0, 0)
}

func (self *InputDispatcher) mouseButtonCallback(button, state int) {
	log.Println("Mouse Button pressed", button, state)
}

func (self *InputDispatcher) mouseWheelCallback(position int) {
	log.Println("Mouse wheel", position)
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
