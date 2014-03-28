package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Visual_Type(t *testing.T) {
	visual := Visual{}
	assert.Equal(t, VISUAL, visual.Type())
}

func Test_GetVisual(t *testing.T) {
	visual := Visual{}
	holder := &TestHolder{}
	holder.AddComponent(&visual)

	assert.Equal(t, &visual, GetVisual(holder))
}
