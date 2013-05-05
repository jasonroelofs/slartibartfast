package math3d

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_Perspective(t *testing.T) {
	matrix := Perspective(90, 1, 0, 1)

	expected := Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, -1,
		0, 0, -0, 0,
	}

	assert.Equal(t, expected, matrix)
}

func Test_LookAt(t *testing.T) {
	matrix := LookAt(Vector{1, 0, 0}, Vector{0, 0, 0}, Vector{0, 1, 0})

	expected := Matrix{
		0, 0, 0, 0,
		0, 1, 0, 0,
		-1, 0, 0, 0,
		0, 0, -1, 1,
	}

	assert.Equal(t, expected, matrix)
}
