package math3d

import (
	"math"
)

// Vector is your standard 3D vector.
// Homogeneous coords is handled when needed
type Vector struct {
	X, Y, Z float32
}

func ZeroVector() Vector {
	return Vector{0, 0, 0}
}

func (self Vector) Add(other Vector) Vector {
	return Vector{
		self.X + other.X,
		self.Y + other.Y,
		self.Z + other.Z,
	}
}

func (self Vector) Sub(other Vector) Vector {
	return Vector{
		self.X - other.X,
		self.Y - other.Y,
		self.Z - other.Z,
	}
}

func (self Vector) Times(value float32) Vector {
	return Vector{
		self.X * value,
		self.Y * value,
		self.Z * value,
	}
}

func (self Vector) Length() float32 {
	return float32(math.Sqrt(float64(self.X * self.X + self.Y * self.Y + self.Z * self.Z)))
}

func (self Vector) Normalize() Vector {
	length := self.Length()
	return Vector {
		self.X / length,
		self.Y / length,
		self.Z / length,
	}
}

func (self Vector) Dot(other Vector) float32 {
	return self.X * other.X + self.Y * other.Y + self.Z * other.Z;
}

func (self Vector) Cross(other Vector) Vector {
	return Vector{
		self.Y * other.Z - self.Z * other.Y,
		self.Z * other.X - self.X * other.Z,
		self.X * other.Y - self.Y * other.X,
	}
}
