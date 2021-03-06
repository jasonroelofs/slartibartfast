package behaviors

import (
	"components"
	"core"
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

// SetUpEntity :: EntityListener
func (self *Transform) SetUpEntity(entity *core.Entity) {
}

// TearDownEntity :: EntityListener
func (self *Transform) TearDownEntity(entity *core.Entity) {
}

func (self *Transform) Update(deltaT float32) {
	var transform *components.Transform

	for _, entity := range self.entitySet.Entities() {
		transform = components.GetTransform(entity)

		if transform.UsingPositionOf == nil {
			self.processMovement(deltaT, transform)
			self.processRotation(deltaT, transform)
		} else {
			transform.Position = components.GetTransform(transform.UsingPositionOf).Position
		}
	}
}

func (self *Transform) processMovement(deltaT float32, component *components.Transform) {
	moveDir := component.MoveDir()

	if component.MoveRelativeToRotation {
		moveDir = component.Rotation.Inverse().TimesV(moveDir)
	}

	moveDir = moveDir.Times(component.Speed).Scale(deltaT)
	component.Position = component.Position.Add(moveDir)
}

func (self *Transform) processRotation(deltaT float32, component *components.Transform) {
	rotateDir := component.RotateDir().Times(component.RotationSpeed).Scale(deltaT)

	if rotateDir.X != 0 {
		component.CurrentPitch += rotateDir.X
	}

	if rotateDir.Y != 0 {
		component.CurrentYaw += rotateDir.Y
	}

	if rotateDir.Z != 0 {
		component.CurrentRoll += rotateDir.Z
	}

	component.RecalculateCurrentRotation()
}
