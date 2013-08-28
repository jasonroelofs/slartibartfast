package main

import (
	"components"
	"events"
)

// InputMapping for Fixed top-down camera controls.
// Very similar to FixedY but responds to the Panning and Zooming events,
// allowing different key mappings to this movement.
var FixedCameraMapping components.InputEventMap
var FixedCameraInput *components.Input

func init() {
	FixedCameraMapping = components.InputEventMap{
		events.PanUp:    panUp,
		events.PanDown:  panDown,
		events.PanLeft:  panLeft,
		events.PanRight: panRight,
		events.ZoomOut:  zoomOut,
		events.ZoomIn:   zoomIn,
	}

	FixedCameraInput = &components.Input{
		Mapping: FixedCameraMapping,
		Polling: []events.EventType{
			events.PanUp,
			events.PanDown,
			events.PanLeft,
			events.PanRight,
			events.ZoomOut,
			events.ZoomIn,
		},
	}
}

func panUp(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingForward(event.Pressed)
}

func panDown(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingBackward(event.Pressed)
}

func panLeft(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingLeft(event.Pressed)
}

func panRight(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingRight(event.Pressed)
}

func zoomIn(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingDown(event.Pressed)
}

func zoomOut(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingUp(event.Pressed)
}
