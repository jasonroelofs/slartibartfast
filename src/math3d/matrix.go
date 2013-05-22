package math3d

// The default type of Matrix used in any 3D situation is the 4x4 matrix
// So the default "Matrix" type will be that. If I need a different size one,
// I'll use the explicit names like Matrix3.
//
// For reference, here's the layout:
//
//		 0,  1,  2,  3,
//     4,  5,  6,  7,
//     8,  9, 10, 11,
//    12, 13, 14, 15
//
type Matrix [16]float32

func IdentityMatrix() (matrix Matrix) {
	matrix[0] = 1
	matrix[5] = 1
	matrix[10] = 1
	matrix[15] = 1
	return
}

func (self Matrix) Times(other Matrix) Matrix {
	return Matrix{
		self[0] * other[0] + self[4] * other[1] + self[8] * other[2] + self[12] * other[3],
		self[1] * other[0] + self[5] * other[1] + self[9] * other[2] + self[13] * other[3],
		self[2] * other[0] + self[6] * other[1] + self[10] * other[2] + self[14] * other[3],
		self[3] * other[0] + self[7] * other[1] + self[11] * other[2] + self[15] * other[3],

		self[0] * other[4] + self[4] * other[5] + self[8] * other[6] + self[12] * other[7],
		self[1] * other[4] + self[5] * other[5] + self[9] * other[6] + self[13] * other[7],
		self[2] * other[4] + self[6] * other[5] + self[10] * other[6] + self[14] * other[7],
		self[3] * other[4] + self[7] * other[5] + self[11] * other[6] + self[15] * other[7],

		self[0] * other[8] + self[4] * other[9] + self[8] * other[10] + self[12] * other[11],
		self[1] * other[8] + self[5] * other[9] + self[9] * other[10] + self[13] * other[11],
		self[2] * other[8] + self[6] * other[9] + self[10] * other[10] + self[14] * other[11],
		self[3] * other[8] + self[7] * other[9] + self[11] * other[10] + self[15] * other[11],

		self[0] * other[12] + self[4] * other[13] + self[8] * other[14] + self[12] * other[15],
		self[1] * other[12] + self[5] * other[13] + self[9] * other[14] + self[13] * other[15],
		self[2] * other[12] + self[6] * other[13] + self[10] * other[14] + self[14] * other[15],
		self[3] * other[12] + self[7] * other[13] + self[11] * other[14] + self[15] * other[15],
	}
}

func (self Matrix) Transpose() Matrix {
	return Matrix{
		self[0], self[4], self[8], self[12],
		self[1], self[5], self[9], self[13],
		self[2], self[6], self[10], self[14],
		self[3], self[7], self[11], self[15],
	}
}
