package math3d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PositionMatrix(t *testing.T) {
	vector := Vector{1, 2, 3}

	expected := Matrix {
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}

	assert.Equal(t, expected, PositionMatrix(vector))
}

func Test_ScaleMatrix(t *testing.T) {
	scale := Vector{1, 2, 3}

	expected := Matrix {
		1, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 3, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, expected, ScaleMatrix(scale))
}

func Test_RotationMatrix(t *testing.T) {
	rotation := NewQuaternion()
	expected := Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, expected, RotationMatrix(rotation))
}
