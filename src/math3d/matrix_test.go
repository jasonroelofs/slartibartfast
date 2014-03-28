package math3d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DefaultMatrixStartsOutZeros(t *testing.T) {
	var matrix Matrix

	for i := 0; i < 16; i++ {
		assert.Equal(t, 0, matrix[i])
	}
}

func Test_IdentityMatrix_ReturnsIdentityMatrix(t *testing.T) {
	matrix := IdentityMatrix()
	assert.Equal(t, 1, matrix[0])
	assert.Equal(t, 0, matrix[1])
	assert.Equal(t, 0, matrix[2])
	assert.Equal(t, 0, matrix[3])

	assert.Equal(t, 0, matrix[4])
	assert.Equal(t, 1, matrix[5])
	assert.Equal(t, 0, matrix[6])
	assert.Equal(t, 0, matrix[7])

	assert.Equal(t, 0, matrix[8])
	assert.Equal(t, 0, matrix[9])
	assert.Equal(t, 1, matrix[10])
	assert.Equal(t, 0, matrix[11])

	assert.Equal(t, 0, matrix[12])
	assert.Equal(t, 0, matrix[13])
	assert.Equal(t, 0, matrix[14])
	assert.Equal(t, 1, matrix[15])
}

func Test_Times_MultipliesMatricies(t *testing.T) {
	matrix1 := Matrix{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}

	matrix2 := Matrix{
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
		1, 1, 1, 1,
	}

	expects := Matrix{
		4, 4, 4, 4,
		4, 4, 4, 4,
		4, 4, 4, 4,
		4, 4, 4, 4,
	}

	got := matrix1.Times(matrix2)

	assert.Equal(t, expects, got)
}

func Test_Matrix_Transpose(t *testing.T) {
	matrix := Matrix{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}

	expected := Matrix{
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16,
	}

	assert.Equal(t, expected, matrix.Transpose())
}

func Benchmark_Multiplcation(b *testing.B) {
	matrix1 := Matrix{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}

	matrix2 := Matrix{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16,
	}

	for i := 0; i < b.N; i++ {
		matrix1.Times(matrix2)
	}
}
