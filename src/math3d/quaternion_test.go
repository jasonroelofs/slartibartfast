package math3d

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewQuaternion(t *testing.T) {
	quat := NewQuaternion()
	assert.Equal(t, Quaternion{1, 0, 0, 0}, quat)
}

func Test_QuatFromAngleAxis(t *testing.T) {
	zero := NewQuaternion()
	quat := QuatFromAngleAxis(90, Vector{0, 0, 0})

	assert.NotEqual(t, zero, quat)
}

func Test_QuatFromAxes(t *testing.T) {
	quat := QuatFromAxes(Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1})
	assert.Equal(t, NewQuaternion(), quat)
}

func Test_QuatFromRotationMatrix(t *testing.T) {
	matrix := IdentityMatrix()

	assert.Equal(t, NewQuaternion(), QuatFromRotationMatrix(matrix))
}

func Test_QuatFromRotationMatrix_FirstElse(t *testing.T) {
	matrix := Matrix{
		1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, Quaternion{0, 1, 0, 0}, QuatFromRotationMatrix(matrix))
}

func Test_QuatFromRotationMatrix_SecondElse(t *testing.T) {
	matrix := Matrix{
		-1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, Quaternion{0, 0, 1, 0}, QuatFromRotationMatrix(matrix))
}

func Test_QuatFromRotationMatrix_ThirdElse(t *testing.T) {
	matrix := Matrix{
		-1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, Quaternion{0, 0, 0, 1}, QuatFromRotationMatrix(matrix))
}

func Test_Quaternion_Times(t *testing.T) {
	quat1 := Quaternion{1, 1, 2, 3}
	quat2 := Quaternion{1, 1, 2, 3}

	assert.Equal(t, Quaternion{-13, 2, 4, 6}, quat2.Times(quat1))
}

func Test_Quaternion_TimesV(t *testing.T) {
	quat := Quaternion{0, 0, 1, 0}
	vec := Vector{1, 0, 0}

	assert.Equal(t, Vector{-1, 0, 0}, quat.TimesV(vec))
}

func Test_Quaternion_Normalize(t *testing.T) {
	quat := Quaternion{1080, 10, 10, 10}

	assert.True(t, quat.Length() > 1.0)
	assert.True(t, quat.Normalize().Length() <= 1.0)
}
