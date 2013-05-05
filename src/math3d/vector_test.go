package math3d

import (
	"github.com/stretchrcom/testify/assert"
	"math"
	"testing"
)

func Test_VectorDefaultsZero(t *testing.T) {
	vector := Vector{}
	assert.Equal(t, 0, vector.X)
	assert.Equal(t, 0, vector.Y)
	assert.Equal(t, 0, vector.Z)
}

func Test_Add(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{4, 5, 6}

	assert.Equal(t,  Vector{5, 7, 9}, v1.Add(v2))
}

func Test_Sub(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{4, 5, 6}

	assert.Equal(t,  Vector{3, 3, 3}, v2.Sub(v1))
}

func Test_Times(t *testing.T) {
	vec := Vector{1, 2, 3}

	assert.Equal(t,  Vector{10, 20, 30}, vec.Times(10))
}

func Test_Length(t *testing.T) {
	vec := Vector{2, 3, 4}

	assert.Equal(t, float32(math.Sqrt(29)), vec.Length())
}

func Test_Normalize(t *testing.T) {
	vec := Vector{2, 0, 0}

	assert.Equal(t, Vector{1, 0, 0}, vec.Normalize())
}

func Test_Dot(t *testing.T) {
	v1 := Vector{1, 0, 0}
	v2 := Vector{0, 1, 0}

	assert.Equal(t, 0, v1.Dot(v2))
	assert.Equal(t, 1, v1.Dot(v1))
}

func Test_Cross(t *testing.T) {
	v1 := Vector{1, 0, 0}
	v2 := Vector{0, 1, 0}

	assert.Equal(t, Vector{0, 0, 1}, v1.Cross(v2))
}

