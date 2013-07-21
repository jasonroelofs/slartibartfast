package core

import (
	"components"
	"fmt"
	"math3d"
)

type componentTypeMap map[components.ComponentType]components.Component

type Entity struct {
	// Implements components.ComponentHolder

	Name string

	components       componentTypeMap
	destroyNextFrame bool
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

// Destroy flags this Entity to be destroyed at the beginning
// of the next frame.
func (self *Entity) Destroy() {
	self.destroyNextFrame = true
}

// AddComponent adds a given component to this Entity
// +component+ *must* be a pointer or the system won't work.
func (self *Entity) AddComponent(component components.Component) {
	self.components[component.Type()] = component
}

// GetComponent returns the component on this Entity of the given type
// To get the underlying struct, typeAssert it with .(*components.[component struct])
func (self *Entity) GetComponent(componentType components.ComponentType) components.Component {
	return self.components[componentType]
}

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
