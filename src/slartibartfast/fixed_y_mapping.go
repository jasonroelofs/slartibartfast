package main

import (
	"components"
	"events"
	"log"
	"math3d"
)

// Simple 2d-based movement when looking down at the player from above.
var FixedYMapping components.InputEventMap
var FixedYInput *components.Input

func init() {
	FixedYMapping = components.InputEventMap{
		events.MoveForward:  fixedYMoveForward,
		events.MoveBackward: fixedYMoveBackward,
		events.MoveLeft:     fixedYMoveLeft,
		events.MoveRight:    fixedYMoveRight,
		events.MouseMove:    fixedYMouseMoved,
		events.Primary:      fire,
	}

	FixedYInput = &components.Input{
		Mapping: FixedYMapping,
		Polling: []events.EventType{
			events.MoveForward, events.MoveBackward, events.MoveLeft, events.MoveRight,
		},
	}
}

func fixedYMoveForward(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingForward(event.Pressed)
}

func fixedYMoveBackward(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingBackward(event.Pressed)
}

func fixedYMoveLeft(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingLeft(event.Pressed)
}

func fixedYMoveRight(entity components.ComponentHolder, event events.Event) {
	components.GetTransform(entity).MovingRight(event.Pressed)
}

// Make the entity always look towards the direction of the mouse cursor
// Always rotate around +Y (yaw)
func fixedYMouseMoved(entity components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(entity)

	transform.CurrentYaw = math3d.RadToDeg(
		math3d.Atan2(float32(event.MouseYDiff), float32(event.MouseXDiff)),
	)*-1 + 90
}

func fire(entity components.ComponentHolder, event events.Event) {
	log.Println("Fire zee missiles!")
}
