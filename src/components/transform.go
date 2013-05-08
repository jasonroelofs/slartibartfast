package components

import (
	"math3d"
)

// Transform holds location and rotation data of the holding Entity
type Transform struct {
	Position math3d.Vector
	Scale    math3d.Vector

	// I don't want a constructor method yet I want the default of Scale
	// to be 1, 1, 1. This may or may not work out well.
	scaleInitialized bool
}

func (self Transform) Type() ComponentType {
	return TRANSFORM
}

func GetTransform(holder ComponentHolder) *Transform {
	return holder.GetComponent(TRANSFORM).(*Transform)
}

func (self Transform) TransformMatrix() math3d.Matrix {
	if self.Scale == math3d.ZeroVector() && !self.scaleInitialized {
		self.Scale = math3d.Vector{1, 1, 1}
	}
	self.scaleInitialized = true

	position := math3d.Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		self.Position.X, self.Position.Y, self.Position.Z, 1,
	}

	scale := math3d.Matrix{
		self.Scale.X, 0, 0, 0,
		0, self.Scale.Y, 0, 0,
		0, 0, self.Scale.Z, 0,
		0, 0, 0, 1,
	}

	return position.Times(scale)
}
