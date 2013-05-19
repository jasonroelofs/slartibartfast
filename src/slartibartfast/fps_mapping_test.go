package main

import (
	"components"
	"core"
	"events"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_MoveForward(t *testing.T) {
	entity := core.NewEntity()
	event := events.Event{}

	FPSMapping[events.MoveForward](entity, event)

	transform := components.GetTransform(entity)
	assert.Equal(t, math3d.Vector{0, 0, -1}, transform.Position)
}
