package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Input_Type(t *testing.T) {
	visual := Input{}
	assert.Equal(t, INPUT, visual.Type())
}

func Test_GetInput(t *testing.T) {
	visual := Input{}
	holder := &TestHolder{}
	holder.AddComponent(&visual)

	assert.Equal(t, &visual, GetInput(holder))
}
