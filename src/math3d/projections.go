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

// LookAt calculates the LookAt matrix for the given parameters (camera-centric)
func LookAt(position, lookAt, up Vector) Matrix {
	zAxis := (position.Sub(lookAt)).Normalize()
	xAxis := up.Normalize().Cross(zAxis)
	yAxis := zAxis.Cross(xAxis)

	dotX := xAxis.Dot(position)
	dotY := yAxis.Dot(position)
	dotZ := zAxis.Dot(position)

	viewMatrix := Matrix{
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		-dotX, -dotY, -dotZ, 1,
	}

	return viewMatrix
}

// ViewMatrix calculates a full View Matrix from a Position and Rotation (Quaternion)
// This is the Transpose of the Rotation matrix * Inverse of the position, because
// we are moving the world into our view, not moving our view of the world.
func ViewMatrix(position Vector, rotation Quaternion) Matrix {
	viewMatrix := RotationMatrix(rotation).Transpose()
	translation := position.Scale(-1)

	viewMatrix[12] = translation.X
	viewMatrix[13] = translation.Y
	viewMatrix[14] = translation.Z

	return viewMatrix
}
