package components

import (
	"math3d"
)

// Transform holds location and rotation data of the holding Entity
type Transform struct {
	Position math3d.Vector
	Rotation math3d.Quaternion
	Scale    math3d.Vector

	// Speed at which this Entity will move along each axis
	// Defaults to 1 unit / second
	Speed math3d.Vector

	// Speed at which this entity will rotate around each axis
	// Defaults to 45 degrees / second
	RotationSpeed math3d.Vector

	// Should this Entity move in the direction the entity
	// is currently facing?
	MoveRelativeToRotation bool

	// Can also use Euler values to rotate this transform
	CurrentYaw, CurrentPitch, CurrentRoll float32

	// Should this Transform rotate around a fixed Up axis, and
	// what is that axis?
	FixedUp          bool
	FixedUpDirection math3d.Vector

	// The direction this Entity is currently moving
	moveDirection math3d.Vector

	// The direction this Entity is rotating
	rotateDirection math3d.Vector

	// If set, this Transform will be given the position of this Entity every frame,
	// ignoring any other attempts to move it.
	UsingPositionOf ComponentHolder
}

func NewTransform() Transform {
	return Transform{
		Position:      math3d.Vector{0, 0, 0},
		Scale:         math3d.Vector{1, 1, 1},
		Speed:         math3d.Vector{1, 1, 1},
		Rotation:      math3d.NewQuaternion(),
		RotationSpeed: math3d.Vector{45, 45, 45},

		// Default to +Y as the Up dir
		FixedUp:          true,
		FixedUpDirection: math3d.Vector{0, 1, 0},
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
	// Trying to look at ourselves makes for weird rotations
	if lookAtPoint == self.Position {
		return
	}

	// TODO Calculate the current Up if FixedUp is false

	rotMatrix := math3d.LookAt(self.Position, lookAtPoint, self.FixedUpDirection)
	self.Rotation = math3d.QuatFromRotationMatrix(rotMatrix)
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

// Rotating sets a vector defining the axis and direction (positive / negative)
// of the rotation. So to rotate this entity around the X axis in negative degrees,
// give Vector{-1, 0, 0}. Multiple rotation directions can be running at the
// same time, subsequent calls to this method will combine the directions.
// As with Moving, don't use dir to define Speed, instead use RotationSpeed to
// set how fast this Entity will rotate.
func (self *Transform) Rotating(dir math3d.Vector) {
	self.rotateDirection = self.rotateDirection.Add(dir)
}

// RotateDir returns a normalized Vector from Rotating
func (self *Transform) RotateDir() math3d.Vector {
	return self.rotateDirection.Normalize()
}

// RecalculateCurrentRotation takes the current Euler angles (roll, pitch, yaw)
// and rebuilds the internal Quaternion to match those new values
func (self *Transform) RecalculateCurrentRotation() {
	rollQuat := math3d.QuatFromAngleAxis(self.CurrentRoll, math3d.Vector{0, 0, 1})
	pitchQuat := math3d.QuatFromAngleAxis(self.CurrentPitch, math3d.Vector{1, 0, 0})
	yawQuat := math3d.QuatFromAngleAxis(self.CurrentYaw, math3d.Vector{0, 1, 0})

	self.Rotation = rollQuat.Times(pitchQuat).Times(yawQuat)
}
