package behaviors

import (
	"components"
	"core"
	"github.com/stretchr/testify/assert"
	"math3d"
	"testing"
)

func getTestAnimation() (*Animation, *core.EntityDB) {
	entityDB := core.NewEntityDB()

	behavior := NewAnimation(entityDB)

	return behavior, entityDB
}

func Test_NewAnimation_RegistersWithDB(t *testing.T) {
	behavior, _ := getTestAnimation()
	assert.NotNil(t, behavior.entitySet)
}

func Test_Update_UpdatesAPositionAnimationOverTime(t *testing.T) {
	behavior, entityDb := getTestAnimation()

	entity := core.NewEntity()
	entity.AddComponent(
		components.NewPositionAnimation(math3d.Vector{2, 2, 2}, 1, func() { }),
	)
	entityDb.RegisterEntity(entity)

	// Move half way
	behavior.Update(0.5)

	transform := components.GetTransform(entity)
	assert.Equal(t, math3d.Vector{1, 1, 1}, transform.Position)

	// Move the rest of the way
	behavior.Update(0.5)

	transform = components.GetTransform(entity)
	assert.Equal(t, math3d.Vector{2, 2, 2}, transform.Position)
}

func Test_Update_UpdatesOverMultipleSeconds(t *testing.T) {
	behavior, entityDb := getTestAnimation()

	entity := core.NewEntity()
	entity.AddComponent(
		components.NewPositionAnimation(math3d.Vector{10, 10, 10}, 10, func() { }),
	)
	entityDb.RegisterEntity(entity)

	var i float32 = 1
	for ; i <= 10; i++ {
		behavior.Update(1)

		transform := components.GetTransform(entity)
		assert.Equal(t, math3d.Vector{i, i, i}, transform.Position)
	}
}

func Test_Update_RemovesAFinishedOneTimeAnimation(t *testing.T) {
	behavior, entityDb := getTestAnimation()

	entity := core.NewEntity()
	entity.AddComponent(
		components.NewPositionAnimation(math3d.Vector{10, 10, 10}, 0.5, func() { }),
	)
	entityDb.RegisterEntity(entity)

	behavior.Update(1)

	assert.Nil(t, entity.GetComponent(components.ANIMATION))
}

func Test_Update_CallsCompletionCallbackOnCompletion(t *testing.T) {
	behavior, entityDb := getTestAnimation()

	callbackCalled := false

	entity := core.NewEntity()
	entity.AddComponent(
		components.NewPositionAnimation(math3d.Vector{10, 10, 10}, 0.5, func() { callbackCalled = true }),
	)
	entityDb.RegisterEntity(entity)

	behavior.Update(1)

	assert.True(t, callbackCalled)
}
