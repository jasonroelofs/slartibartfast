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

	graphical := NewTransform(entityDB)

	return graphical, entityDB
}

func Test_NewTransform(t *testing.T) {
	graphical, _ := getTestTransform()
	assert.NotNil(t, graphical.entitySet)
}

func Test_Update_AppliesMovementDirOnTransforms(t *testing.T) {
	transform, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	eTransform := components.GetTransform(entity)
	eTransform.Moving(math3d.Vector{1, 0, 0})

	// No time passed? no change
	transform.Update(0)
	assert.Equal(t, math3d.Vector{0, 0, 0}, eTransform.Position)

	// Time passed? change!
	transform.Update(1)
	assert.Equal(t, math3d.Vector{1, 0, 0}, eTransform.Position)
}

func Test_Update_AppliesSpeedToMovingDir(t *testing.T) {
	transform, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	eTransform := components.GetTransform(entity)
	eTransform.Moving(math3d.Vector{1, 0, 0})
	eTransform.Speed = math3d.Vector{2, 2, 2}

	// Time passed? change!
	transform.Update(0.5)
	assert.Equal(t, math3d.Vector{1, 0, 0}, eTransform.Position)
}

func Test_Update_AppliesRotationDir(t *testing.T) {
	transform, entityDb := getTestTransform()

	entity := core.NewEntity()
	entityDb.RegisterEntity(entity)

	eTransform := components.GetTransform(entity)
	eTransform.Rotating(math3d.Vector{0, 1, 0})

	startingQuat := eTransform.Rotation

	// No time passed? no change
	transform.Update(0)
	assert.Equal(t, startingQuat, eTransform.Rotation)

	// Time passed? change!
	transform.Update(1)
	assert.NotEqual(t, startingQuat, eTransform.Rotation)
}
