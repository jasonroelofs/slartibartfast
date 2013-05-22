package core

import (
	"components"
	"math3d"
)

// A Camera defines how one looks at the current scene
type Camera struct {
	Up         math3d.Vector

	projection math3d.Matrix
	lookAt     math3d.Vector

	// Cameras link to an Entity which is their physical
	// presence in the scene
	Entity     *Entity
}

func NewCamera() *Camera {
	camera := new(Camera)
	camera.Up = math3d.Vector{0, 1, 0}
	camera.Entity = NewEntity()
	camera.Entity.Name = "Camera"

	return camera
}

func (self Camera) Position() math3d.Vector {
	return components.GetTransform(self.Entity).Position
}

func (self *Camera) SetPosition(newPosition math3d.Vector) {
	components.GetTransform(self.Entity).Position = newPosition
}

func (self *Camera) SetSpeed(newSpeed math3d.Vector) {
	components.GetTransform(self.Entity).Speed = newSpeed
}

func (self *Camera) AddComponent(component components.Component) {
	self.Entity.AddComponent(component)
}

func (self *Camera) LookAt(point math3d.Vector) {
	self.lookAt = point
}

func (self *Camera) Perspective(fov, aspectRatio, nearPlane, farPlane float32) {
	self.projection = math3d.Perspective(fov, aspectRatio, nearPlane, farPlane)
}

func (self *Camera) ProjectionMatrix() math3d.Matrix {
	return self.projection
}

func (self *Camera) ViewMatrix() math3d.Matrix {
	return math3d.LookAt(
		self.Position(),
		self.lookAt,
		self.Up,
	)
}
