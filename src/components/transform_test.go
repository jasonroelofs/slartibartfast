package components

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_Type(t *testing.T) {
	transform := Transform{}
	assert.Equal(t, TRANSFORM, transform.Type())
}

type TestHolder struct {
	Holding Component
}

func (self *TestHolder) AddComponent(c Component) {
	self.Holding = c
}

func (self *TestHolder) GetComponent(t ComponentType) Component {
	return self.Holding
}

func Test_GetTransform(t *testing.T) {
	transform := Transform{}
	holder := &TestHolder{}
	holder.AddComponent(&transform)

	assert.Equal(t, &transform, GetTransform(holder))
}
