package main

import (
	"components"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Level_Generate_ReturnsEntity(t *testing.T) {
	level := NewLevel()
	entity := level.Generate()

	assert.NotNil(t, entity)
	assert.NotNil(t, components.GetVisual(entity), "Did not generate a visual component")
}
