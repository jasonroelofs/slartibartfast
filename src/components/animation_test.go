package components

import (
	"github.com/stretchr/testify/assert"
	"math3d"
	"testing"
)

func Test_Animation_Type(t *testing.T) {
	anim := Animation{}
	assert.Equal(t, ANIMATION, anim.Type())
}

func Test_GetAnimation(t *testing.T) {
	anim := Animation{}
	holder := &TestHolder{}
	holder.AddComponent(&anim)

	assert.Equal(t, &anim, GetAnimation(holder))
}

func Test_NewPositionAnimation(t *testing.T) {
	moveTo := math3d.Vector{1, 2, 3}
	finished := false
	anim := NewPositionAnimation(moveTo, 1, func() { finished = true })

	assert.Equal(t, "Position", anim.PropertyName)
	assert.Equal(t, moveTo, anim.TweenTo)
	assert.Equal(t, 1, anim.TweenTime)

	anim.CompletionCallback()
	assert.True(t, finished)
}

func Test_Finish_CallsCallbacksIfExists(t *testing.T) {
	moveTo := math3d.Vector{1, 2, 3}
	finished := false
	anim := NewPositionAnimation(moveTo, 1, func() { })

	anim.Finish()
	assert.False(t, finished)

	anim.CompletionCallback = func() { finished = true }

	anim.Finish()
	assert.True(t, finished)
}
