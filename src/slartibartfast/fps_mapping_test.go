package main

import (
	"components"
	"core"
	"events"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

var cardinalTests = []struct {
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

	for _, testValue := range cardinalTests {
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

var turnTests = []struct {
	event       events.EventType
	expectedDir math3d.Vector
}{
	{events.TurnLeft, math3d.Vector{0, -1, 0}},
	{events.TurnRight, math3d.Vector{0, 1, 0}},
}

func Test_Turning(t *testing.T) {
	event := events.Event{}

	for _, testValue := range turnTests {
		entity := core.NewEntity()
		transform := components.GetTransform(entity)

		// Set on pressed
		event.Pressed = true

		FPSMapping[testValue.event](entity, event)
		assert.Equal(t, testValue.expectedDir, transform.RotateDir())

		// And undo the change on key release
		event.Pressed = false

		FPSMapping[testValue.event](entity, event)
		assert.Equal(t, math3d.Vector{0, 0, 0}, transform.RotateDir())
	}
}

func Test_MouseMove(t *testing.T) {
	event := events.Event{
		MouseXDiff: 30,
		MouseYDiff: 40,
	}

	entity := core.NewEntity()
	transform := components.GetTransform(entity)

	FPSMapping[events.MouseMove](entity, event)

	assert.Equal(t, 30, transform.CurrentYaw)
	assert.Equal(t, 40, transform.CurrentPitch)
}

func Test_MouseMove_ConstrainsYaw(t *testing.T) {
	event := events.Event{
		MouseXDiff: 40,
		MouseYDiff: 0,
	}

	entity := core.NewEntity()
	transform := components.GetTransform(entity)
	transform.CurrentYaw = 350

	FPSMapping[events.MouseMove](entity, event)

	assert.Equal(t, 30, transform.CurrentYaw)
}

func Test_MouseMove_ClampsPitch(t *testing.T) {
	event := events.Event{
		MouseXDiff: 0,
		MouseYDiff: 40,
	}

	entity := core.NewEntity()
	transform := components.GetTransform(entity)
	transform.CurrentPitch = 80

	FPSMapping[events.MouseMove](entity, event)

	assert.Equal(t, 89, transform.CurrentPitch)
}
