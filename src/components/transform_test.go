package components

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, 180, transform.CurrentRoll)
	assert.Equal(t, 0, transform.CurrentYaw * -1) // 0 != -0 /shrug
	assert.Equal(t, 180, transform.CurrentPitch)
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

func Test_MovingForward_HandlesMinusZMoveDir(t *testing.T) {
	transform := NewTransform()

	transform.MovingForward(true)
	assert.Equal(t, math3d.Vector{0, 0, -1}, transform.MoveDir())

	transform.MovingForward(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
}

func Test_MovingBackward_HandlesPlusZMoveDir(t *testing.T) {
	transform := NewTransform()

	transform.MovingBackward(true)
	assert.Equal(t, math3d.Vector{0, 0, 1}, transform.MoveDir())

	transform.MovingBackward(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
}

func Test_MovingLeft_HandlesMinusXMoveDir(t *testing.T) {
	transform := NewTransform()

	transform.MovingLeft(true)
	assert.Equal(t, math3d.Vector{-1, 0, 0}, transform.MoveDir())

	transform.MovingLeft(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
}

func Test_MovingRight_HandlesPlusXMoveDir(t *testing.T) {
	transform := NewTransform()

	transform.MovingRight(true)
	assert.Equal(t, math3d.Vector{1, 0, 0}, transform.MoveDir())

	transform.MovingRight(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
}

func Test_MovingUp_HandlesPlusYMoveDir(t *testing.T) {
	transform := NewTransform()

	transform.MovingUp(true)
	assert.Equal(t, math3d.Vector{0, 1, 0}, transform.MoveDir())

	transform.MovingUp(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
}

func Test_MovingDown_HandlesMinusYMoveDir(t *testing.T) {
	transform := NewTransform()

	transform.MovingDown(true)
	assert.Equal(t, math3d.Vector{0, -1, 0}, transform.MoveDir())

	transform.MovingDown(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.MoveDir())
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

func Test_TurningLeft_SetsRotatingMinusY(t *testing.T) {
	transform := NewTransform()

	transform.TurningLeft(true)
	assert.Equal(t, math3d.Vector{0, -1, 0}, transform.rotateDirection)

	transform.TurningLeft(false)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.rotateDirection)
}

func Test_TurningRight_SetsRotatingPositiveY(t *testing.T) {
	transform := NewTransform()

	transform.TurningRight(true)
	assert.Equal(t, math3d.Vector{0, 1, 0}, transform.rotateDirection)

	transform.TurningRight(false)
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
