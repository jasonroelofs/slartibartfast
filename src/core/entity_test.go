package core

import (
	"components"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

// NewEntity()
func Test_ReturnsANewEntity(t *testing.T) {
	entity := NewEntity()
	assert.NotNil(t, entity)
}

func Test_InitializesEntityWithTransformComponent(t *testing.T) {
	entity := NewEntity()

	assert.Equal(t, 1, len(entity.components))
	assert.IsType(t, components.Transform{}, entity.components[0])
}

// AddComponent
func Test_CanAddAComponentToEntity(t *testing.T) {
	entity := NewEntity()
	entity.AddComponent(components.Visual{})

	assert.Equal(t, 2, len(entity.components))
}

// ComponentMap()
func Test_CanReturnBitwiseMapOfComponents(t *testing.T) {
	entity := NewEntity()
	assert.Equal(t, 1, entity.ComponentMap())

	entity.AddComponent(components.Visual{})

	// 01(transform) + 10(visual) == 11
	assert.Equal(t, 3, entity.ComponentMap())
}
