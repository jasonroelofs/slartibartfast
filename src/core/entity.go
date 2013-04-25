package core

import (
	"components"
)

type Entity struct {
	components []components.Component
}

func NewEntity() (entity *Entity) {
	entity = new(Entity)
	entity.AddComponent(components.Transform{})

	return
}

func (self *Entity) AddComponent(component components.Component) {
	self.components = append(self.components, component)
}

func (self *Entity) ComponentMap() (typeMap components.ComponentType) {
	typeMap = 0
	for _, component := range self.components {
		typeMap |= component.Type()
	}
	return
}
