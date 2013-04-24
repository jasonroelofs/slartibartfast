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
type TestComponent struct {
}

func Test_CanAddAComponentToEntity(t *testing.T) {
	entity := NewEntity()
	entity.AddComponent(TestComponent{})

	assert.Equal(t, 2, len(entity.components))
}
