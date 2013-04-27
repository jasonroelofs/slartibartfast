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

func (self *Entity) AddComponent(component components.Component) {
	self.components[component.Type()] = component
}

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
