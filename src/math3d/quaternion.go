package math3d

import "math"

// The (in)famous Quaternion!
// Formulas pulled from http://www.cprogramming.com/tutorial/3d/quaternions.html
type Quaternion struct {
	W, X, Y, Z float32
}

// NewQuaternion returns the Unit Quaternion
func NewQuaternion() Quaternion {
	return Quaternion{1, 0, 0, 0}
}

// QuatFromAngleAxis creates a new Quaterion calcualted from the angle and the
// axis on which the rotation is happening.
func QuatFromAngleAxis(angleInDegrees float32, axis Vector) Quaternion {
	angleRad := (float32(math.Pi) * angleInDegrees) / 180.0
	cosAngle := float32(math.Cos(float64(angleRad / 2)))
	sinAngle := float32(math.Sin(float64(angleRad / 2)))

	return Quaternion{
		cosAngle,
		axis.X * sinAngle,
		axis.Y * sinAngle,
		axis.Z * sinAngle,
	}
}

// Times returns a new Quaternion that is the product of this Quaternion and other
func (self Quaternion) Times(other Quaternion) Quaternion {
	return Quaternion{
		self.W * other.W - self.X * other.X - self.Y * other.Y - self.Z * other.Z,
		self.W * other.X + self.X * other.W + self.Y * other.Z - self.Z * other.Y,
		self.W * other.Y + self.Y * other.W + self.Z * other.X - self.X * other.Z,
		self.W * other.Z + self.Z * other.W + self.X * other.Y - self.Y * other.X,
	}
}

// Length returns the length, or magnitude, of this Quaternion
func (self Quaternion) Length() float32 {
	return float32(math.Sqrt(
		float64(self.W*self.W + self.X*self.X + self.Y*self.Y + self.Z*self.Z)))
}

// Normalize returns a new Quaternion that is the normalized version of the
// current Quaternion
func (self Quaternion) Normalize() Quaternion {
	magnitude := self.Length()

	return Quaternion{
		self.W / magnitude,
		self.X / magnitude,
		self.Y / magnitude,
		self.Z / magnitude,
	}
}

// RotateX returns a new Quaternion that is the current one rotated
// by the given degrees around the X axis
func (self Quaternion) RotateX(degrees float32) Quaternion {
	toApply := QuatFromAngleAxis(degrees, Vector{1, 0, 0})
	return toApply.Times(self)
}

// RotateY returns a new Quaternion that is the current one rotated
// by the given degrees around the Y axis
func (self Quaternion) RotateY(degrees float32) Quaternion {
	toApply := QuatFromAngleAxis(degrees, Vector{0, 1, 0})
	return toApply.Times(self)
}

// RotateZ returns a new Quaternion that is the current one rotated
// by the given degrees around the Z axis
func (self Quaternion) RotateZ(degrees float32) Quaternion {
	toApply := QuatFromAngleAxis(degrees, Vector{0, 0, 1})
	return toApply.Times(self)
}
