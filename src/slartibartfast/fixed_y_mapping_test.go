package main

import (
	"components"
	"core"
	"events"
	"github.com/stretchr/testify/assert"
	"math3d"
	"testing"
)

var fixedYCardinalTests = []struct {
	event       events.EventType
	expectedDir math3d.Vector
}{
	{events.MoveForward, math3d.Vector{0, 0, -1}},
	{events.MoveBackward, math3d.Vector{0, 0, 1}},
	{events.MoveLeft, math3d.Vector{-1, 0, 0}},
	{events.MoveRight, math3d.Vector{1, 0, 0}},
}

func Test_FixedY_CardinalMovement(t *testing.T) {
	event := events.Event{}

	for _, testValue := range fixedYCardinalTests {
		entity := core.NewEntity()
		transform := components.GetTransform(entity)

		// Set on pressed
		event.Pressed = true

		FixedYMapping[testValue.event](entity, event)
		assert.Equal(t, testValue.expectedDir, transform.MoveDir())

		// And undo the change on key release
		event.Pressed = false

		FixedYMapping[testValue.event](entity, event)
		assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
	}
}
