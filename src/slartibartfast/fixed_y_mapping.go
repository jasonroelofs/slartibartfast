package main

import (
	"components"
	"events"
	"math3d"
)

// This is similar to FixedCameraMapping but because we're looking down at an Entity
// and expecting to move the Entity itself we have to change the axes in question.
var FixedYMapping components.InputEventMap

func init() {
	FixedYMapping = components.InputEventMap{
		events.MoveForward:  fixedYMoveForward,
		events.MoveBackward: fixedYMoveBackward,
		events.MoveLeft:     fixedYMoveLeft,
		events.MoveRight:    fixedYMoveRight,
	}
}

func fixedYMoveForward(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{0, 0, -1})
	} else {
		transform.Moving(math3d.Vector{0, 0, 1})
	}
}

func fixedYMoveBackward(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{0, 0, 1})
	} else {
		transform.Moving(math3d.Vector{0, 0, -1})
	}
}

func fixedYMoveLeft(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{-1, 0, 0})
	} else {
		transform.Moving(math3d.Vector{1, 0, 0})
	}
}

func fixedYMoveRight(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	if event.Pressed {
		transform.Moving(math3d.Vector{1, 0, 0})
	} else {
		transform.Moving(math3d.Vector{-1, 0, 0})
	}
}
