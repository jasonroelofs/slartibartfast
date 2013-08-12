package core

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_Append_AddsEntityPointer(t *testing.T) {
	es := NewEntitySet()
	entity := NewEntity()

	es.Append(entity)
	assert.Equal(t, entity, es.entities[0])
}

func Test_Append_DoesNotAddDuplicates(t *testing.T) {
	es := NewEntitySet()
	entity := NewEntity()

	es.Append(entity)
	es.Append(entity)
	es.Append(entity)

	assert.Equal(t, 1, es.Len())
}

func Test_Append_ReturnsWhetherEntityWasAdded(t *testing.T) {
	es := NewEntitySet()
	entity := NewEntity()

	assert.True(t, es.Append(entity), "Successful inclusion should return true")
	assert.False(t, es.Append(entity), "Duplicates should return false")
}

func Test_Len_ReturnsNumberOfEntities(t *testing.T) {
	es := NewEntitySet()
	assert.Equal(t, 0, es.Len())

	es.Append(NewEntity())
	assert.Equal(t, 1, es.Len())
}

func Test_Get_ReturnsEntityOfTheGivenId(t *testing.T) {
	es := NewEntitySet()
	entity := NewEntity()
	entity.Id = 4
	es.Append(entity)

	assert.Equal(t, entity, es.Get(4))
}

func Test_Entities_ReturnsListOfEntities(t *testing.T) {
	es := NewEntitySet()
	entity := NewEntity()
	es.Append(entity)


	assert.Equal(t, []*Entity{entity}, es.Entities())
}

func Test_Delete_RemovesEntityFromList(t *testing.T) {
	es := NewEntitySet()
	entity1 := NewEntity()
	entity1.Id = 10
	entity2 := NewEntity()
	entity2.Id = 20
	entity3 := NewEntity()
	entity3.Id = 30

	es.Append(entity1)
	es.Append(entity2)
	es.Append(entity3)

	es.Delete(entity2)

	assert.Equal(t, 2, len(es.Entities()))
	assert.Nil(t, es.Get(20), "Found entity 2 when it should have gone away")
}

func Test_Contains_ChecksForEntityInList(t *testing.T) {
	es := NewEntitySet()
	entity1 := NewEntity()
	entity1.Id = 10
	entity2 := NewEntity()
	entity2.Id = 20
	entity3 := NewEntity()
	entity3.Id = 30

	es.Append(entity1)
	es.Append(entity2)

	assert.True(t, es.Contains(entity1), "Set didn't contain entity 1?")
	assert.True(t, es.Contains(entity2), "Set didn't contain entity 2?")
	assert.False(t, es.Contains(entity3), "Set contains entity 3?")
}
