package main

import (
	"components"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Player_HasEntity(t *testing.T) {
	player := NewPlayer()

	assert.NotNil(t, player)
	assert.NotNil(t, components.GetVisual(player.GetEntity()), "Did not generate a visual component")
}

func Test_Player_HasInput(t *testing.T) {
	player := NewPlayer()
	assert.NotNil(t, components.GetInput(player.GetEntity()), "No input component found")
}
