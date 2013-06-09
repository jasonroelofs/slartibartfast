package behaviors

import (
	"components"
	"core"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"testing"
)

func getTestTransform() (*Transform, *core.EntityDB) {
	entityDB := new(core.EntityDB)

	behavior := NewTransform(entityDB)

	return behavior, entityDB
}

func Test_NewTransform(t *testing.T) {
	behavior, _ := getTestTransform()
	assert.NotNil(t, behavior.entitySet)
}

func Test_Update_AppliesMovementDirOnTransforms(t *testing.T) {
	behavior, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	transform := components.GetTransform(entity)
	transform.Moving(math3d.Vector{1, 0, 0})

	// No time passed? no change
	behavior.Update(0)
	assert.Equal(t, math3d.Vector{0, 0, 0}, transform.Position)

	// Time passed? change!
	behavior.Update(1)
	assert.Equal(t, math3d.Vector{1, 0, 0}, transform.Position)
}

func Test_Update_AppliesSpeedToMovingDir(t *testing.T) {
	behavior, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	transform := components.GetTransform(entity)
	transform.Moving(math3d.Vector{1, 0, 0})
	transform.Speed = math3d.Vector{2, 2, 2}

	// Time passed? change!
	behavior.Update(0.5)
	assert.Equal(t, math3d.Vector{1, 0, 0}, transform.Position)
}

func Test_Update_MovesWithRotationIfSoFlagged(t *testing.T) {
	behavior, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	transform := components.GetTransform(entity)
	transform.Moving(math3d.Vector{1, 0, 0})
	transform.Speed = math3d.Vector{1, 1, 1}
	transform.Rotation = math3d.Quaternion{0, 0, 1, 0}
	transform.MoveRelativeToRotation = true

	// Time passed? change!
	behavior.Update(1)
	assert.Equal(t, math3d.Vector{-1, 0, 0}, transform.Position)
}

func Test_Update_AppliesRotationDir(t *testing.T) {
	behavior, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	transform := components.GetTransform(entity)
	transform.Rotating(math3d.Vector{0, 1, 0})

	startingQuat := transform.Rotation

	// No time passed? no change
	behavior.Update(0)
	assert.Equal(t, startingQuat, transform.Rotation)

	// Time passed? change!
	behavior.Update(1)
	assert.NotEqual(t, startingQuat, transform.Rotation)
}

func Test_Update_AppliesEulerAngles(t *testing.T) {
	behavior, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	transform := components.GetTransform(entity)
	startingQuat := transform.Rotation
	transform.CurrentPitch = 45

	behavior.Update(1)
	assert.NotEqual(t, startingQuat, transform.Rotation)
}

func Test_Update_CopiesPositionFromLinkedEntity(t *testing.T) {
	behavior, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	followee := core.NewEntity()
	followeeTransform := components.GetTransform(followee)
	followeeTransform.Position = math3d.Vector{10, 11, -12}

	transform := components.GetTransform(entity)
	transform.UsingPositionOf = followee

	behavior.Update(1)

	assert.Equal(t, math3d.Vector{10, 11, -12}, transform.Position)
}
