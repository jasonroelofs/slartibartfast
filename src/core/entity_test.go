package core

import (
	"components"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_NewEntity_ReturnsANewEntity(t *testing.T) {
	entity := NewEntity()
	assert.NotNil(t, entity)
}

func Test_NewEntity_InitializesEntityWithTransformComponent(t *testing.T) {
	entity := NewEntity()

	assert.IsType(t, &components.Transform{}, entity.components[components.TRANSFORM])
}

func Test_NewEntityAt_TakesAStartingPosition(t *testing.T) {
	entity := NewEntityAt(math3d.Vector{1, 2, 3})
	transform := components.GetTransform(entity)
	assert.Equal(t, math3d.Vector{1, 2, 3}, transform.Position)
}

func Test_Destroy_FlagsEntityForDeletion(t *testing.T) {
	entity := NewEntity()
	assert.False(t, entity.destroyNextFrame)

	entity.Destroy()

	assert.True(t, entity.destroyNextFrame)
}

func Test_AddComponent(t *testing.T) {
	entity := NewEntity()
	visual := new(components.Visual)
	entity.AddComponent(visual)

	assert.Equal(t, visual, entity.components[components.VISUAL])
}

func Test_GetComponent(t *testing.T) {
	entity := NewEntity()
	visual := new(components.Visual)
	entity.AddComponent(visual)

	assert.Equal(t, visual, entity.GetComponent(components.VISUAL))
}

func Test_ComponentMap_ReturnsBitwiseMapOfComponents(t *testing.T) {
	entity := NewEntity()
	assert.Equal(t, 1, entity.ComponentMap())

	entity.AddComponent(new(components.Visual))

	// 01(transform) + 10(visual) == 11
	assert.Equal(t, 3, entity.ComponentMap())
}
