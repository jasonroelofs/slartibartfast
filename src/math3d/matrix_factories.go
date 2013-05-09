package math3d

func PositionMatrix(position Vector) Matrix {
	return Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		position.X, position.Y, position.Z, 1,
	}
}

func ScaleMatrix(scale Vector) Matrix {
	return Matrix{
		scale.X, 0, 0, 0,
		0, scale.Y, 0, 0,
		0, 0, scale.Z, 0,
		0, 0, 0, 1,
	}
}

func RotationMatrix(rotation Quaternion) Matrix {
	rotXSq := (rotation.X * rotation.X)
	rotYSq := (rotation.Y * rotation.Y)
	rotZSq := (rotation.Z * rotation.Z)
	w := rotation.W
	x := rotation.X
	y := rotation.Y
	z := rotation.Z

	return Matrix{
		1 - 2*rotYSq - 2*rotZSq, 2*x*y - 2*w*z, 2*x*z + 2*w*y, 0,
		2*x*y + 2*w*z, 1 - 2*rotXSq - 2*rotZSq, 2*y*z + 2*w*x, 0,
		2*x*z - 2*w*y, 2*y*z - 2*w*x, 1 - 2*rotXSq - 2*rotYSq, 0,
		0, 0, 0, 1,
	}
}
