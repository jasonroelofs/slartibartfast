package main

import (
	"components"
	"events"
)

// InputMapping for FPS controls
var FPSMapping components.InputEventMap

func init() {
	FPSMapping = components.InputEventMap{
		events.MoveForward:  moveForward,
		events.MoveBackward: moveBackward,
		events.MoveLeft:     moveLeft,
		events.MoveRight:    moveRight,
	}
}

// Move the given entity in the -Z direction
func moveForward(entity components.ComponentHolder, event events.Event) {
	if event.Pressed {
		transform := components.GetTransform(entity)
		transform.Position.Z -= 1
		// transform.Moving(math3d.Vector{0, 0, -1})
	}
}

// Move the given entity in the +Z direction
func moveBackward(entity components.ComponentHolder, event events.Event) {
	if event.Pressed {
		transform := components.GetTransform(entity)
		transform.Position.Z += 1
	}
}

// Move the given entity in the -X direction
func moveLeft(entity components.ComponentHolder, event events.Event) {
	if event.Pressed {
		transform := components.GetTransform(entity)
		transform.Position.X -= 1
	}
}

// Move the given entity in the +X direction
func moveRight(entity components.ComponentHolder, event events.Event) {
	if event.Pressed {
		transform := components.GetTransform(entity)
		transform.Position.X += 1
	}
}
