package components

import (
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func Test_NewTransform_InitializesGoodDefaults(t *testing.T) {
	transform := NewTransform()
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.Position)
	assert.Equal(t, math3d.Vector{1, 1, 1}, transform.Scale)
	assert.Equal(t, math3d.Quaternion{1, 0, 0, 0}, transform.Rotation)
	assert.Equal(t, math3d.Vector{45, 45, 45}, transform.RotationSpeed)
	assert.True(t, transform.FixedUp)
	assert.Equal(t, math3d.Vector{0, 1, 0}, transform.FixedUpDirection)
}

func Test_Transform_Type(t *testing.T) {
	transform := NewTransform()
	assert.Equal(t, TRANSFORM, transform.Type())
}

func Test_GetTransform(t *testing.T) {
	transform := NewTransform()
	holder := &TestHolder{}
	holder.AddComponent(&transform)

	assert.Equal(t, &transform, GetTransform(holder))
}

func Test_LookAt_ChangesRotationToLookAtPoint(t *testing.T) {
	transform := NewTransform()
	transform.LookAt(math3d.Vector{0, 0, 10})

	assert.Equal(t, math3d.Quaternion{0, 0, 1, 0}, transform.Rotation)
}

func Test_LookAt_DoesNothingIfLookAtIsPosition(t *testing.T) {
	transform := NewTransform()
	transform.LookAt(math3d.Vector{0, 0, 0})

	assert.Equal(t, math3d.NewQuaternion(), transform.Rotation)
}

func Test_Transform_TransformMatrix_DefaultsIdentity(t *testing.T) {
	transform := NewTransform()

	assert.Equal(t, math3d.IdentityMatrix(), transform.TransformMatrix())
}

func Test_Transform_TransformMatrix_AppliesPositionTransformation(t *testing.T) {
	transform := NewTransform()
	transform.Position = math3d.Vector{1, 2, 3}

	expected := math3d.Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		1, 2, 3, 1,
	}

	assert.Equal(t, expected, transform.TransformMatrix())
}

func Test_Transform_TransformMatrix_AppliesScaleTransformation(t *testing.T) {
	transform := NewTransform()
	transform.Scale = math3d.Vector{2, 3, 4}

	expected := math3d.Matrix{
		2, 0, 0, 0,
		0, 3, 0, 0,
		0, 0, 4, 0,
		0, 0, 0, 1,
	}

	assert.Equal(t, expected, transform.TransformMatrix())
}

func Test_Transform_TransformMatrix_AppliesRotationTransformation(t *testing.T) {
	quat := math3d.Quaternion{1, 1, 2, 3}

	transform := Transform{
		Position: math3d.Vector{0, 0, 0},
		Scale: math3d.Vector{1, 1, 1},
		Rotation: quat,
	}

	notExpected := math3d.Matrix{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	assert.NotEqual(t, notExpected, transform.TransformMatrix())
}

func Test_Moving_SetsMoveDirection(t *testing.T) {
	transform := NewTransform()

	transform.Moving(math3d.Vector{1, 0, 0})
	assert.Equal(t, math3d.Vector{1, 0, 0}, transform.moveDirection)

	transform.Moving(math3d.Vector{0, 0, 1})
	assert.Equal(t, math3d.Vector{1, 0, 1}, transform.moveDirection)

	transform.Moving(math3d.Vector{-1, 0, 0})
	assert.Equal(t, math3d.Vector{0, 0, 1}, transform.moveDirection)

	transform.Moving(math3d.Vector{0, 0, -1})
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.moveDirection)
}

func Test_MoveDir_ReturnsNormalizedMoveDirection(t *testing.T) {
	transform := NewTransform()

	transform.Moving(math3d.Vector{1, 0, 0})
	transform.Moving(math3d.Vector{0, 0, 1})
	transform.Moving(math3d.Vector{0, 1, 0})

	assert.True(t, (1 - transform.MoveDir().Length()) < 0.0001)
}

func Test_Rotating_SetsRotateDirection(t *testing.T) {
	transform := NewTransform()

	transform.Rotating(math3d.Vector{1, 0, 0})
	assert.Equal(t, math3d.Vector{1, 0, 0}, transform.rotateDirection)

	transform.Rotating(math3d.Vector{0, 0, 1})
	assert.Equal(t, math3d.Vector{1, 0, 1}, transform.rotateDirection)

	transform.Rotating(math3d.Vector{-1, 0, 0})
	assert.Equal(t, math3d.Vector{0, 0, 1}, transform.rotateDirection)

	transform.Rotating(math3d.Vector{0, 0, -1})
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.rotateDirection)
}

func Test_RotatingDir_ReturnsNormalizedDirection(t *testing.T) {
	transform := NewTransform()

	transform.Rotating(math3d.Vector{1, 0, 0})
	transform.Rotating(math3d.Vector{0, 0, 1})
	transform.Rotating(math3d.Vector{0, 1, 0})

	assert.True(t, (1 - transform.RotateDir().Length()) < 0.0001)
}

func Test_Halt_StopsAllMovement(t *testing.T) {
	transform := NewTransform()
	transform.Moving(math3d.Vector{1, 0, 0})

	transform.Halt()

	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.moveDirection)
}

func Test_Halt_StopsAllRotations(t *testing.T) {
	transform := NewTransform()
	transform.Rotating(math3d.Vector{1, 0, 0})

	transform.Halt()

	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.rotateDirection)
}
