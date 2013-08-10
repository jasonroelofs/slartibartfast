package main

import (
	"components"
	"core"
	"events"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

var panningCardinalTests = []struct {
	event       events.EventType
	expectedDir math3d.Vector
}{
	{events.PanUp, math3d.Vector{0, 1, 0}},
	{events.PanDown, math3d.Vector{0, -1, 0}},
	{events.PanLeft, math3d.Vector{-1, 0, 0}},
	{events.PanRight, math3d.Vector{1, 0, 0}},
}

func Test_FixedCamera_CardinalMovement(t *testing.T) {
	event := events.Event{}

	for _, testValue := range panningCardinalTests {
		entity := core.NewEntity()
		transform := components.GetTransform(entity)

		// Set on pressed
		event.Pressed = true

		FixedCameraMapping[testValue.event](entity, event)
		assert.Equal(t, testValue.expectedDir, transform.MoveDir())

		// And undo the change on key release
		event.Pressed = false

		FixedCameraMapping[testValue.event](entity, event)
		assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
	}
}
