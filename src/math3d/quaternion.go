package math3d

import "math"

// The (in)famous Quaternion!
// Formulas pulled from http://www.cprogramming.com/tutorial/3d/quaternions.html
type Quaternion struct {
	W, X, Y, Z float32
}

func NewQuaternion() Quaternion {
	return Quaternion{1, 0, 0, 0}
}

func QuatFromAngleAxis(angle float32, axis Vector) Quaternion {
	angleRad := (float32(math.Pi) * angle) / 180.0
	cosAngle := float32(math.Cos(float64(angleRad / 2)))
	sinAngle := float32(math.Sin(float64(angleRad / 2)))

	return Quaternion{
		cosAngle,
		axis.X * sinAngle,
		axis.Y * sinAngle,
		axis.Z * sinAngle,
	}
}

func (self Quaternion) Times(other Quaternion) Quaternion {
	return Quaternion{
		self.W * other.W - self.X * other.X - self.Y * other.Y - self.Z * other.Z,
		self.W * other.X + self.X * other.W + self.Y * other.Z - self.Z * other.Y,
		self.W * other.Y - self.X * other.Z + self.Y * other.W + self.Z * other.X,
		self.W * other.Z + self.X * other.Y - self.Y * other.X + self.Z * other.W,
	}
}

func (self Quaternion) Length() float32 {
	return float32(math.Sqrt(
		float64(self.W*self.W + self.X*self.X + self.Y*self.Y + self.Z*self.Z)))
}

func (self Quaternion) Normalize() Quaternion {
	magnitude := self.Length()

	return Quaternion{
		self.W / magnitude,
		self.X / magnitude,
		self.Y / magnitude,
		self.Z / magnitude,
	}
}
