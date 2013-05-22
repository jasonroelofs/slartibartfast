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
	Entity    *Entity
	transform *components.Transform
}

func NewCamera() *Camera {
	camera := new(Camera)
	camera.Entity = NewEntity()
	camera.Entity.Name = "Camera"

	camera.transform = components.GetTransform(camera.Entity)

	return camera
}

//
// components.Transform pass-through
// A camera contains an Entity to keep track if its position
// and orientation in the world, and to support hooking into
// other behaviors as needed, like input.
//

func (self Camera) Position() math3d.Vector {
	return self.transform.Position
}

func (self *Camera) SetPosition(newPosition math3d.Vector) {
	self.transform.Position = newPosition
}

func (self *Camera) Rotation() math3d.Quaternion {
	return self.transform.Rotation
}

func (self *Camera) SetRotation(rotation math3d.Quaternion) {
	self.transform.Rotation = rotation
}

func (self *Camera) SetSpeed(newSpeed math3d.Vector) {
	self.transform.Speed = newSpeed
}

func (self *Camera) AddComponent(component components.Component) {
	self.Entity.AddComponent(component)
}

//
// Projection and View Matrix methods
//

// LookAt changes the camera's orientation so that it is looking in
// the direction of the requested point.
func (self *Camera) LookAt(point math3d.Vector) {
	self.transform.LookAt(point)
}

// Perspective calcualtes the perspective matrix this camera should apply
// when rendering a view.
func (self *Camera) Perspective(fov, aspectRatio, nearPlane, farPlane float32) {
	self.projection = math3d.Perspective(fov, aspectRatio, nearPlane, farPlane)
}

// ProjectionMatrix returns the calcualted projection matrix
func (self *Camera) ProjectionMatrix() math3d.Matrix {
	return self.projection
}

// ViewMatrix calculates and returns the current view matrix
func (self *Camera) ViewMatrix() math3d.Matrix {
	return math3d.ViewMatrix(self.Position(), self.Rotation())
}
