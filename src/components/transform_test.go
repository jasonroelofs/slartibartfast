package components

import (
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_Transform_Type(t *testing.T) {
	transform := Transform{}
	assert.Equal(t, TRANSFORM, transform.Type())
}

func Test_GetTransform(t *testing.T) {
	transform := Transform{}
	holder := &TestHolder{}
	holder.AddComponent(&transform)

	assert.Equal(t, &transform, GetTransform(holder))
}

func Test_Transform_TransformMatrix_DefaultsIdentity(t *testing.T) {
	transform := Transform{}

	assert.Equal(t, math3d.IdentityMatrix(), transform.TransformMatrix())
}

func Test_Transform_TransformMatrix_AppliesPositionTransformation(t *testing.T) {
	transform := Transform{
		Position: math3d.Vector{1, 2, 3},
	}

	expected := math3d.Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}

	assert.Equal(t, expected, transform.TransformMatrix())
}

func Test_Transform_TransformMatrix_AppliesScaleTransformation(t *testing.T) {
	transform := Transform{
		Position: math3d.Vector{0, 0, 0},
		Scale: math3d.Vector{2, 3, 4},
	}

	expected := math3d.Matrix{
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, 4, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, expected, transform.TransformMatrix())
}
