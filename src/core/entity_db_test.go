package core

import (
	"components"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

// Construction
func Test_StartsWithEmptyListOfEntities(t *testing.T) {
	db := EntityDB{}
	assert.Equal(t, 0, len(db.allEntities))
}

// Entity Management
func Test_KeepsTrackOfEntitiesAndComponents(t *testing.T) {
	db := EntityDB{}
	entity := NewEntity()
	entity.AddComponent(new(components.Visual))

	db.RegisterEntity(entity)

	assert.Equal(t, 1, len(db.allEntities))
}

//
// Entity Listeners
//

type TestListener struct {
	setUpEntities []*Entity
}

func (self *TestListener) SetUpEntity(entity *Entity) {
	self.setUpEntities = append(self.setUpEntities, entity)
}

// Registering a Listener
func Test_EntityListenersCanRegisterWithDB(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entitySet := db.RegisterListener(listener, components.TRANSFORM, components.VISUAL)

	assert.Equal(t, 0, len(entitySet.Entities))
	assert.Equal(t, 1, len(db.listeners))
}

// SetupEntity callback to listeners
func Test_ListenerNotifiedOfNewEntityMatchingComponents(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entity := NewEntity()

	db.RegisterListener(listener, components.TRANSFORM)
	db.RegisterEntity(entity)

	assert.Equal(t, 1, len(listener.setUpEntities))
	assert.Equal(t, entity, listener.setUpEntities[0])
}

func Test_ListenerNotNotifiedOfNewEntityIfComponentsDontMatch(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entity := Entity{}

	db.RegisterListener(listener, components.TRANSFORM)
	db.RegisterEntity(&entity)

	assert.Equal(t, 0, len(listener.setUpEntities))
}

func Test_AddsEntityToListenerEntityListIfComponentsMatch(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entity := NewEntity()

	entitySet := db.RegisterListener(listener, components.TRANSFORM)
	db.RegisterEntity(entity)

	assert.Equal(t, 1, len(entitySet.Entities))
	assert.Equal(t, entity, entitySet.Entities[0])
}

func Test_RegisterEntity_ProperlyWorksAgainstMultipleComponentTypes(t *testing.T) {
	db := EntityDB{}

	listener1 := new(TestListener)
	listener2 := new(TestListener)
	listener3 := new(TestListener)

	es1 := db.RegisterListener(listener1, components.TRANSFORM)
	es2 := db.RegisterListener(listener2, components.TRANSFORM, components.VISUAL)
	es3 := db.RegisterListener(listener3, components.TRANSFORM, components.INPUT)

	entity := NewEntity() // Has transform by default
	entity.AddComponent(new(components.Input))

	db.RegisterEntity(entity)

	// 1 and 3 match this entity
	assert.Equal(t, entity, es1.Entities[0])
	assert.Equal(t, entity, es3.Entities[0])

	// But 2 does not
	assert.Equal(t, 0, len(es2.Entities))
}

// CleanUpEntity callback to listeners

// Do all Entities need pointers back to the db that created them?
// so Entity can inform DB when components change?
