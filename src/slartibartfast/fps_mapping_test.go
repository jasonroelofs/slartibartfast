package main

import (
	"components"
	"core"
	"events"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

var eventTests = []struct {
	event       events.EventType
	expectedDir math3d.Vector
}{
	{events.MoveForward, math3d.Vector{0, 0, -1}},
	{events.MoveBackward, math3d.Vector{0, 0, 1}},
	{events.MoveRight, math3d.Vector{1, 0, 0}},
	{events.MoveLeft, math3d.Vector{-1, 0, 0}},
}

func Test_CardinalMovement(t *testing.T) {
	event := events.Event{}

	for _, testValue := range eventTests {
		entity := core.NewEntity()
		transform := components.GetTransform(entity)

		// Set on pressed
		event.Pressed = true

		FPSMapping[testValue.event](entity, event)
		assert.Equal(t, testValue.expectedDir, transform.MoveDir())

		// And undo the change on key release
		event.Pressed = false

		FPSMapping[testValue.event](entity, event)
		assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
	}
}
