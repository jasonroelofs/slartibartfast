package behaviors

import (
	"components"
	"core"
	"math3d"
)

// The Transform behavior takes care of moving Entities around
// the scene as necessary
type Transform struct {
	entitySet *core.EntitySet
}

func NewTransform(entityDB *core.EntityDB) *Transform {
	transform := new(Transform)
	transform.entitySet = entityDB.RegisterListener(transform, components.TRANSFORM)
	return transform
}

// EntityListener
func (self *Transform) SetUpEntity(entity *core.Entity) {
}

func (self *Transform) Update(deltaT float32) {
	var transform *components.Transform
	var moveDir math3d.Vector

	for _, entity := range self.entitySet.Entities {
		transform = components.GetTransform(entity)
		moveDir = transform.
			MoveDir().
			Times(transform.Speed).
			Scale(deltaT)

		transform.Position = transform.Position.Add(moveDir)
	}
}
