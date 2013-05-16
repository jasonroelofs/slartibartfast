package math3d

// Perspective calculates the Perspective Matrix for the given parameters
func Perspective(fov, aspectRatio, nearPlane, farPlane float32) (matrix Matrix) {
	fovRad := DegToRad(fov)
	f := 1.0 / Tan(fovRad/2)

	matrix[0] = f / aspectRatio
	matrix[5] = f
	matrix[10] = (nearPlane + farPlane) / (nearPlane - farPlane)
	matrix[11] = -1
	matrix[14] = (2 * farPlane * nearPlane) / (nearPlane - farPlane)

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
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		-dotX, -dotY, -dotZ, 1,
	}
}
