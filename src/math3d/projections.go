package math3d

import (
	"math"
)

// Perspective calculates the Perspective Matrix for the given parameters
func Perspective(fov, aspectRatio, nearPlane, farPlane float64) (matrix Matrix) {
	fovRad := (fov * math.Pi) / 180.0
	f := 1.0 / math.Tan(fovRad/2)

	matrix[0] = float32(f / aspectRatio)
	matrix[5] = float32(f)
	matrix[10] = float32((nearPlane + farPlane) / (nearPlane - farPlane))
	matrix[11] = -1
	matrix[14] = float32((2 * farPlane * nearPlane) / (nearPlane - farPlane))

	return
}

// LookAt calculates the View Matrix for the given parameters (camera-centric)
func LookAt(position, lookAt, up Vector) Matrix {
	zAxis := (position.Sub(lookAt)).Normalize()
	xAxis := up.Normalize().Cross(zAxis)
	yAxis := zAxis.Cross(xAxis)

	dotX := xAxis.Dot(position)
	dotY := yAxis.Dot(position)
	dotZ := zAxis.Dot(position)

	return Matrix{
		xAxis.X, yAxis.X, zAxis.Z, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		-dotX, -dotY, -dotZ, 1,
	}
}
