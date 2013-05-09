package components

import (
	"math3d"
)

// Transform holds location and rotation data of the holding Entity
type Transform struct {
	Position math3d.Vector
	Scale    math3d.Vector
	Rotation math3d.Quaternion
}

func NewTransform() Transform {
	return Transform{
		Position: math3d.Vector{0, 0, 0},
		Scale:    math3d.Vector{1, 1, 1},
		Rotation: math3d.NewQuaternion(),
	}
}

func (self Transform) Type() ComponentType {
	return TRANSFORM
}

func GetTransform(holder ComponentHolder) *Transform {
	return holder.GetComponent(TRANSFORM).(*Transform)
}

func (self Transform) TransformMatrix() math3d.Matrix {
	position := math3d.PositionMatrix(self.Position)
	scale := math3d.ScaleMatrix(self.Scale)
	rotation := math3d.RotationMatrix(self.Rotation)

	return position.Times(rotation).Times(scale)
}
