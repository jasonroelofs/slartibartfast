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

func (e *Entity) AddComponent(component components.Component) {
	e.components = append(e.components, component)
}
