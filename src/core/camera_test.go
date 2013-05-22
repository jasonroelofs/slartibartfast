package core

import (
	"components"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_NewCamera_SetsGoodDefaults(t *testing.T) {
	camera := NewCamera()
	assert.Equal(t, math3d.Vector{0, 1, 0}, camera.Up)
	assert.Equal(t, math3d.Vector{0, 0, 0}, camera.lookAt)
	assert.Equal(t, math3d.Vector{0, 0, 0}, camera.Position())
	assert.Equal(t, math3d.Matrix{}, camera.ProjectionMatrix())
}

func Test_CanReceiveComponents(t *testing.T) {
	camera := NewCamera()

	input := new(components.Input)
	camera.AddComponent(input)

	assert.Equal(t, input, components.GetInput(camera.Entity))
}

func Test_SetPosition(t *testing.T) {
	camera := NewCamera()
	camera.SetPosition(math3d.Vector{1, 2, 3})

	assert.Equal(t, math3d.Vector{1, 2, 3}, components.GetTransform(camera.Entity).Position)
}

func Test_SetSpeed(t *testing.T) {
	camera := NewCamera()
	camera.SetSpeed(math3d.Vector{1, 2, 3})

	assert.Equal(t, math3d.Vector{1, 2, 3}, components.GetTransform(camera.Entity).Speed)
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

	assert.Equal(t, expected, camera.ProjectionMatrix())
}

func Test_ViewMatrix_CalculatesViewMatrix(t *testing.T) {
	camera := NewCamera()
	camera.Up = math3d.Vector{0, 1, 0}
	camera.LookAt(math3d.Vector{0, 0, 0})
	camera.SetPosition(math3d.Vector{1, 0, 0})

	expected := math3d.Matrix{
		0, 0, 1, 0,
		0, 1, 0, 0,
		-1, 0, 0, 0,
		0, 0, -1, 1,
	}

	assert.Equal(t, expected, camera.ViewMatrix())
}
