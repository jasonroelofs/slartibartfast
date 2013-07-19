package core

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_Append_AddsEntityPointer(t *testing.T) {
	es := EntitySet{}
	entity := NewEntity()

	es.Append(entity)
	assert.Equal(t, entity, es.entities[0])
}

func Test_Len_ReturnsNumberOfEntities(t *testing.T) {
	es := EntitySet{}
	assert.Equal(t, 0, es.Len())

	es.Append(NewEntity())
	assert.Equal(t, 1, es.Len())
}

func Test_Get_ReturnsEntityAtIndex(t *testing.T) {
	es := EntitySet{}
	entity := NewEntity()
	es.Append(entity)

	assert.Equal(t, entity, es.Get(0))
}

func Test_Entities_ReturnsListOfEntities(t *testing.T) {
	es := EntitySet{}
	entity := NewEntity()
	es.Append(entity)


	assert.Equal(t, []*Entity{entity}, es.Entities())
}
