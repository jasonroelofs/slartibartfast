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
	entity := Entity{}
	entity.AddComponent(components.Transform{})
	entity.AddComponent(components.Visual{})

	db.RegisterEntity(&entity)

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

// CleanUpEntity callback to listeners

// Do all Entities need pointers back to the db that created them?
// so Entity can inform DB when components change?
