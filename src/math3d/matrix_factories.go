package math3d

// ScaleMatrix returns the Matrix form of the current Position Vector
func PositionMatrix(position Vector) Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		position.X, position.Y, position.Z, 1,
	}
}

// ScaleMatrix returns the Matrix form of the current Scaling Vector
func ScaleMatrix(scale Vector) Matrix {
	return Matrix{
		scale.X, 0, 0, 0,
		0, scale.Y, 0, 0,
		0, 0, scale.Z, 0,
		0, 0, 0, 1,
	}
}

// RotationMatrix calculates the Matrix form of the given Quaternion
func RotationMatrix(rotation Quaternion) Matrix {
	rotation = rotation.Normalize()
	w := rotation.W
	x := rotation.X
	y := rotation.Y
	z := rotation.Z

	return Matrix{
		1 - (2*y*y + 2*z*z), 2*x*y - 2*w*z, 2*x*z + 2*w*y, 0,
		2*x*y + 2*w*z, 1 - (2*x*x + 2*z*z), 2*y*z - 2*w*x, 0,
		2*x*z - 2*w*y, 2*y*z + 2*w*x, 1 - (2*x*x + 2*y*y), 0,
		0, 0, 0, 1,
	}
}
