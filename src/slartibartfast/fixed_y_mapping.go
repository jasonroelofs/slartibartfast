package main

import (
	"components"
	"events"
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
