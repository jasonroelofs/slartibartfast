package core

import (
	"components"
	"fmt"
	"math3d"
)

type componentTypeMap map[components.ComponentType]components.Component

type Entity struct {
	// Implements components.ComponentHolder

	Id   int
	Name string

	components       componentTypeMap
	destroyNextFrame bool

	// The EntityDB this Entity is registered with
	entityDB EntityDatabase
}

// NewEntity sets up a new Entity for use in the app complete with
// a basic Transform component at the origin
func NewEntity() (entity *Entity) {
	entity = new(Entity)
	entity.components = make(componentTypeMap)

	transform := components.NewTransform()
	entity.AddComponent(&transform)

	return
}

// NewEntityAt sets up a new Entity for use in the app and sets the
// initial Transform component's Position to the given Vector
func NewEntityAt(startingPosition math3d.Vector) *Entity {
	entity := NewEntity()
	transform := components.GetTransform(entity)
	transform.Position = startingPosition

	return entity
}

// Destroy this Entity, removing it and all of it's components entirely
func (self *Entity) Destroy() {
	if self.entityDB != nil {
		self.entityDB.EntityDestroyed(self)
	}
}

// AddComponent adds a given component to this Entity
// +component+ *must* be a pointer or the system won't work.
func (self *Entity) AddComponent(component components.Component) {
	self.components[component.Type()] = component

	if self.entityDB != nil {
		self.entityDB.ComponentAdded(self)
	}
}

// RemoveComponent removes the component of the given ComponentType from
// this Entity.
func (self *Entity) RemoveComponent(componentType components.ComponentType) (removed components.Component) {
	removed = self.components[componentType]

	// Run callbacks before deletion because tear-down might need the component
	// being cleared out
	if self.entityDB != nil {
		self.entityDB.ComponentRemoved(self, removed)
	}

	delete(self.components, componentType)
	return
}

// GetComponent returns the component on this Entity of the given type
// To get the underlying struct, typeAssert it with .(*components.[component struct])
func (self *Entity) GetComponent(componentType components.ComponentType) components.Component {
	return self.components[componentType]
}

// ComponentMap returns the bitmap representation of the components this Entity currently uses
func (self *Entity) ComponentMap() (typeMap components.ComponentType) {
	typeMap = 0
	for componentType, _ := range self.components {
		typeMap |= componentType
	}
	return
}

func (self *Entity) String() string {
	return fmt.Sprintf("Entity \"%s\" :: ComponentMap => %v", self.Name, self.components)
}
