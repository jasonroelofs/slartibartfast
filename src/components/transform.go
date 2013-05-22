package components

import (
	"math3d"
)

// Transform holds location and rotation data of the holding Entity
type Transform struct {
	Position math3d.Vector
	Rotation math3d.Quaternion
	Scale    math3d.Vector
	Speed    math3d.Vector

	// The direction this Entity is currently moving
	moveDirection math3d.Vector
}

func NewTransform() Transform {
	return Transform{
		Position: math3d.Vector{0, 0, 0},
		Scale:    math3d.Vector{1, 1, 1},
		Speed:    math3d.Vector{1, 1, 1},
		Rotation: math3d.NewQuaternion(),
	}
}

func (self Transform) Type() ComponentType {
	return TRANSFORM
}

func GetTransform(holder ComponentHolder) *Transform {
	return holder.GetComponent(TRANSFORM).(*Transform)
}

// LookAt changes this Transform's Rotation such that it's facing the point
// For now, always assumes fixed Y-up axis
func (self *Transform) LookAt(lookAtPoint math3d.Vector) {
	fixedUp := math3d.Vector{0, 1, 0}

	forward := self.Position.Sub(lookAtPoint).Normalize()
	right := fixedUp.Cross(forward).Normalize()
	up := forward.Cross(right).Normalize()

	self.Rotation = math3d.QuatFromAxes(right, up, forward)
}

// TransformMatrix calculates and returns the full transformation matrix
// for this Transform, combining Scale, Rotation, and Position.
func (self Transform) TransformMatrix() math3d.Matrix {
	position := math3d.PositionMatrix(self.Position)
	scale := math3d.ScaleMatrix(self.Scale)
	rotation := math3d.RotationMatrix(self.Rotation)

	return position.Times(rotation).Times(scale)
}

// Moving sets the current direction this Transform should be moving
// Subsequent calls to this method will apply the vector to the existing
// move direction. This gets normalized when requested via MoveDir so set
// the expected rate of change via Speed
func (self *Transform) Moving(dir math3d.Vector) {
	self.moveDirection = self.moveDirection.Add(dir)
}

// MoveDir normalizes and returns the current direction in which this
// transform is moving.
func (self *Transform) MoveDir() math3d.Vector {
	return self.moveDirection.Normalize()
}
