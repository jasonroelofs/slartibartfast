package core

import (
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_NewCamera_SetsGoodDefaults(t *testing.T) {
	camera := NewCamera()
	assert.Equal(t, math3d.Vector{0, 1, 0}, camera.Up)
	assert.Equal(t, math3d.Vector{0, 0, 0}, camera.LookAt)
	assert.Equal(t, math3d.Vector{0, 0, 0}, camera.Position)
	assert.Equal(t, math3d.Matrix{}, camera.Projection)
}

func Test_Perspective_SetsPerspectiveMatrixAsProjection(t *testing.T) {
	camera := NewCamera()
	camera.Perspective(90.0, 1, 0.0, 1.0)

	expected := math3d.Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, -1,
		0, 0, -0, 0,
	}

	assert.Equal(t, expected, camera.Projection)
}

func Test_ViewMatrix_CalculatesViewMatrix(t *testing.T) {
	camera := NewCamera()
	camera.Position = math3d.Vector{1, 0, 0}
	camera.LookAt = math3d.Vector{0, 0, 0}
	camera.Up = math3d.Vector{0, 1, 0}

	expected := math3d.Matrix{
		0, 0, 0, 0,
		0, 1, 0, 0,
		-1, 0, 0, 0,
		0, 0, -1, 1,
	}

	assert.Equal(t, expected, camera.ViewMatrix())
}