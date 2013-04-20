package input

import (
	"events"
	"fmt"
	"github.com/go-gl/glfw"
)

type callbackMap map[events.EventType]func(events.Event)
type keyMap map[int]events.EventType

type InputMapping struct {
	callbacks   callbackMap
	keyMappings keyMap
}

func NewInput() *InputMapping {
	mapping := InputMapping{
		callbacks:   make(callbackMap),
		keyMappings: make(keyMap),
	}

	// Set up testing key mappings
	mapping.keyMappings['Q'] = events.QUIT

	glfw.SetKeyCallback(mapping.keyCallback)

	return &mapping
}

// On registers a callback to be called in the occurance of an event of type EventType.
// The callback will include event details, including key hit, and whether the key was
// pressed or released
func (mapper *InputMapping) On(event events.EventType, callback func(events.Event)) {
	mapper.callbacks[event] = callback
}

func (mapper *InputMapping) keyCallback(key, state int) {
	fmt.Println("Key pressed! ", key, state)
	keyToEvent := mapper.keyMappings[key]
	if keyToEvent != events.NULL {
		eventToCallback := mapper.callbacks[keyToEvent]

		if eventToCallback != nil {
			event := events.Event{}
			eventToCallback(event)
		}
	}
}
