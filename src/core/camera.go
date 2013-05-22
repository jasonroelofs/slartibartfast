package core

import (
	"components"
	"math3d"
)

// A Camera defines how one looks at the current scene
type Camera struct {
	projection math3d.Matrix
	lookAt     math3d.Vector

	// Cameras link to an Entity which is their physical
	// presence in the scene
	Entity *Entity
}

func NewCamera() *Camera {
	camera := new(Camera)
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

func (self *Camera) Rotation() math3d.Quaternion {
	return components.GetTransform(self.Entity).Rotation
}

func (self *Camera) SetRotation(rotation math3d.Quaternion) {
	components.GetTransform(self.Entity).Rotation = rotation
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
	return math3d.ViewMatrix(self.Position(), self.Rotation())
}
