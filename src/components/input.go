package components

import (
	"events"
)

// Entity-based input callbacks.
type EntityEventCallback func(ComponentHolder, events.Event)
type InputEventMap map[events.EventType]EntityEventCallback

// The Input Component handles mapping input events to individual
// Entities. Any number of events, and any number of any type of event,
// can happen at any frame. Also the input handling system, at least with
// any button input (keyboard and mouse) fires events only when state changes,
// e.g. button down, button up. There is no event for "button being held now",
// so the callbacks and resulting code must instead handle a change of state
// instead of making small per-frame changes themselves.
type Input struct {
	// A mapping of Event => Callback on what events
	// this input wants to receive
	Mapping InputEventMap
}

func (self Input) Type() ComponentType {
	return INPUT
}

func GetInput(holder ComponentHolder) *Input {
	return holder.GetComponent(INPUT).(*Input)
}
