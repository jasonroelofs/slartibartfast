package math3d

import (
	"github.com/Jragonmiris/mathgl"
)

type Matrix mathgl.Mat4f

func Perspective(fov, aspectRatio, nearPlane, farPlane float64) Matrix {
	return Matrix(mathgl.Perspective(fov, aspectRatio, nearPlane, farPlane))
}

func IdentityMatrix() Matrix {
	return Matrix(mathgl.Ident4f())
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) Matrix {
	return Matrix(mathgl.LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ))
}

func (self Matrix) Times(other Matrix) Matrix {
	return Matrix(mathgl.Mat4f(self).Mul4(mathgl.Mat4f(other)))
}
