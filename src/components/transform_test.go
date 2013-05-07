package components

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_Transform_Type(t *testing.T) {
	transform := Transform{}
	assert.Equal(t, TRANSFORM, transform.Type())
}

func Test_GetTransform(t *testing.T) {
	transform := Transform{}
	holder := &TestHolder{}
	holder.AddComponent(&transform)

	assert.Equal(t, &transform, GetTransform(holder))
}
