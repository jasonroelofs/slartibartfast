package main

import (
	"components"
	"events"
	"math3d"
)

// InputMapping for FPS controls
var FPSMapping components.InputEventMap

func init() {
	FPSMapping = components.InputEventMap{
		events.MoveForward:  moveForward,
		events.MoveBackward: moveBackward,
		events.MoveLeft:     moveLeft,
		events.MoveRight:    moveRight,
		events.TurnLeft:     turnLeft,
		events.TurnRight:    turnRight,
	}
}

// Move the given entity in the -Z direction
func moveForward(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{0, 0, -1})
	} else {
		transform.Moving(math3d.Vector{0, 0, 1})
	}
}

// Move the given entity in the +Z direction
func moveBackward(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{0, 0, 1})
	} else {
		transform.Moving(math3d.Vector{0, 0, -1})
	}
}

// Move the given entity in the -X direction
func moveLeft(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{-1, 0, 0})
	} else {
		transform.Moving(math3d.Vector{1, 0, 0})
	}
}

// Move the given entity in the +X direction
func moveRight(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{1, 0, 0})
	} else {
		transform.Moving(math3d.Vector{-1, 0, 0})
	}
}

// Rotate entity around its Y axis
func turnLeft(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Rotating(math3d.Vector{0, 1, 0})
	} else {
		transform.Rotating(math3d.Vector{0, -1, 0})
	}
}

func turnRight(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Rotating(math3d.Vector{0, -1, 0})
	} else {
		transform.Rotating(math3d.Vector{0, 1, 0})
	}
}
