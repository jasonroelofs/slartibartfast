package core

import (
	"components"
)

type componentTypeMap map[components.ComponentType]components.Component

type Entity struct {
	components componentTypeMap
}

func NewEntity() (entity *Entity) {
	entity = new(Entity)
	entity.components = make(componentTypeMap)
	entity.AddComponent(new(components.Transform))

	return
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
