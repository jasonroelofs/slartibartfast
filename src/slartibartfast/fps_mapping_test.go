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
	event := events.Event{Pressed: true}

	for _, testValue := range eventTests {
		entity := core.NewEntity()

		FPSMapping[testValue.event](entity, event)

		transform := components.GetTransform(entity)
		assert.Equal(t, testValue.expectedDir, transform.MoveDir())
	}
}
