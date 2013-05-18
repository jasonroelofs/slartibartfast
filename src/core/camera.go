package core

import (
	"components"
	"math3d"
)

// A Camera defines how one looks at the current scene
type Camera struct {
	Projection math3d.Matrix
	LookAt     math3d.Vector
	Up         math3d.Vector

	// Cameras link to an Entity which is their physical
	// presence in the scene
	entity     *Entity
}

func NewCamera() *Camera {
	camera := new(Camera)
	camera.Up = math3d.Vector{0, 1, 0}
	camera.entity = NewEntity()
	return camera
}

func (self Camera) Position() math3d.Vector {
	return components.GetTransform(self.entity).Position
}

func (self *Camera) SetPosition(newPosition math3d.Vector) {
	components.GetTransform(self.entity).Position = newPosition
}

func (self *Camera) Perspective(fov, aspectRatio, nearPlane, farPlane float32) {
	self.Projection = math3d.Perspective(fov, aspectRatio, nearPlane, farPlane)
}

func (self *Camera) ProjectionMatrix() math3d.Matrix {
	return self.Projection
}

func (self *Camera) ViewMatrix() math3d.Matrix {
	return math3d.LookAt(
		self.Position(),
		self.LookAt,
		self.Up,
	)
}
