package core

import (
	"components"
	"github.com/stretchr/testify/assert"
	"math3d"
	"testing"
)

type BogusDB struct {
	// Implements EntityDatabase
	entityWithAddedComponents   []*Entity
	entityWithRemovedComponents []*Entity
	removedComponents           []components.Component
	destroyedEntities           []*Entity
}

func (self *BogusDB) ComponentAdded(entity *Entity) {
	self.entityWithAddedComponents = append(self.entityWithAddedComponents, entity)
}

func (self *BogusDB) ComponentRemoved(entity *Entity, component components.Component) {
	self.entityWithRemovedComponents = append(self.entityWithRemovedComponents, entity)
	self.removedComponents = append(self.removedComponents, component)
}

func (self *BogusDB) EntityDestroyed(entity *Entity) {
	self.destroyedEntities = append(self.destroyedEntities, entity)
}

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

func Test_Destroy_TellsEntityDBToShutDownEntity(t *testing.T) {
	db := new(BogusDB)
	entity := NewEntity()
	entity.entityDB = db

	entity.Destroy()

	assert.Equal(t, entity, db.destroyedEntities[0])
}

func Test_AddComponent(t *testing.T) {
	entity := NewEntity()
	visual := new(components.Visual)
	entity.AddComponent(visual)

	assert.Equal(t, visual, entity.components[components.VISUAL])
}

func Test_AddComponent_TellsEntityDBOfNewComponent(t *testing.T) {
	db := new(BogusDB)
	entity := NewEntity()
	entity.entityDB = db

	entity.AddComponent(new(components.Visual))

	assert.Equal(t, entity, db.entityWithAddedComponents[0])
}

func Test_RemoveComponent(t *testing.T) {
	entity := NewEntity()
	visual := new(components.Visual)
	entity.AddComponent(visual)
	entity.RemoveComponent(components.VISUAL)

	assert.Nil(t, entity.components[components.VISUAL])
}

func Test_RemoveComponent_TellsEntityDBOfRemovedComponent(t *testing.T) {
	db := new(BogusDB)
	entity := NewEntity()
	entity.entityDB = db

	transform := entity.RemoveComponent(components.TRANSFORM)

	assert.Equal(t, entity, db.entityWithRemovedComponents[0])
	assert.Equal(t, transform, db.removedComponents[0])
}

func Test_RemoveComponent_ReturnsTheRemovedComponent(t *testing.T) {
	entity := NewEntity()
	visual := new(components.Visual)
	entity.AddComponent(visual)

	assert.Equal(t, visual, entity.RemoveComponent(components.VISUAL))
}

func Test_RemoveComponent_NoOpsIfNoComponent(t *testing.T) {
	entity := NewEntity()
	entity.RemoveComponent(components.VISUAL)

	assert.Nil(t, entity.components[components.VISUAL])
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
