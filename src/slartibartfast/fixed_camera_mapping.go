package main

import (
	"components"
	"events"
	"math3d"
)

// InputMapping for Fixed top-down camera controls
var FixedCameraMapping components.InputEventMap

func init() {
	FixedCameraMapping = components.InputEventMap{
		events.PanUp:    panForward,
		events.PanDown:  panBackward,
		events.PanLeft:  panLeft,
		events.PanRight: panRight,
	}
}

func panForward(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{0, 1, 0})
	} else {
		transform.Moving(math3d.Vector{0, -1, 0})
	}
}

func panBackward(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{0, -1, 0})
	} else {
		transform.Moving(math3d.Vector{0, 1, 0})
	}
}

func panLeft(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{-1, 0, 0})
	} else {
		transform.Moving(math3d.Vector{1, 0, 0})
	}
}

func panRight(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{1, 0, 0})
	} else {
		transform.Moving(math3d.Vector{-1, 0, 0})
	}
}
