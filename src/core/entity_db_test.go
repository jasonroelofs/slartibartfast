package core

import (
	"components"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

// Construction
func Test_StartsWithEmptyListOfEntities(t *testing.T) {
	db := EntityDB{}
	assert.Equal(t, 0, db.allEntities.Len())
}

// Entity Management
func Test_KeepsTrackOfEntitiesAndComponents(t *testing.T) {
	db := EntityDB{}
	entity := NewEntity()
	entity.AddComponent(new(components.Visual))

	db.RegisterEntity(entity)

	assert.Equal(t, 1, db.allEntities.Len())
}

//
// Entity Listeners
//

type TestListener struct {
	entitySet        *EntitySet
	setUpEntities    []*Entity
	tearDownEntities []*Entity
}

func (self *TestListener) SetUpEntity(entity *Entity) {
	self.setUpEntities = append(self.setUpEntities, entity)
}

func (self *TestListener) TearDownEntity(entity *Entity) {
	self.tearDownEntities = append(self.tearDownEntities, entity)
}

// Registering a Listener
func Test_EntityListenersCanRegisterWithDB(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entitySet := db.RegisterListener(listener, components.TRANSFORM, components.VISUAL)

	assert.Equal(t, 0, entitySet.Len())
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

	assert.Equal(t, 1, entitySet.Len())
	assert.Equal(t, entity, entitySet.Get(0))
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
	assert.Equal(t, entity, es1.Get(0))
	assert.Equal(t, entity, es3.Get(0))

	// But 2 does not
	assert.Equal(t, 0, es2.Len())
}

//
// Dirty Entity Processing
//

func RegisterDBAndListeners() (*EntityDB, [3]*TestListener) {
	db := new(EntityDB)

	listener1 := new(TestListener)
	listener2 := new(TestListener)
	listener3 := new(TestListener)

	listener1.entitySet = db.RegisterListener(listener1, components.TRANSFORM)
	listener2.entitySet = db.RegisterListener(listener2, components.TRANSFORM, components.VISUAL)
	listener3.entitySet = db.RegisterListener(listener3, components.TRANSFORM, components.INPUT)

	return db, [3]*TestListener{listener1, listener2, listener3}
}

func Test_Update_ProcessesDeletedEntities_RemovesEntitiesFromSets(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	db.RegisterEntity(entity)

	entity.Destroy()
	db.Update()

	assert.Equal(t, 0, listeners[0].entitySet.Len())
	assert.Equal(t, 0, listeners[1].entitySet.Len())
	assert.Equal(t, 0, listeners[2].entitySet.Len())
}

func Test_Update_ProcessesDeletedEntities_TellsListenersToCleanUpEntities(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	db.RegisterEntity(entity)

	entity.Destroy()
	db.Update()

	// Triggers on listeners who want the entity
	assert.Equal(t, entity, listeners[0].tearDownEntities[0])
	// Ignores all others
	assert.Equal(t, 0, len(listeners[1].tearDownEntities))
	assert.Equal(t, 0, len(listeners[2].tearDownEntities))
}

func Test_Update_TellsListenersToSetUpNewEntityOnNewComponents(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	db.RegisterEntity(entity)

	entity.AddComponent(new(components.Visual))
	db.Update()

	// Check set up callback triggered
	assert.Equal(t, entity, listeners[1].setUpEntities[0])
	// And added to the listener's entity set
	assert.Equal(t, entity, listeners[1].entitySet.Entities()[0])
}

func Test_Update_TellsListenersToTearDownEntityOnComponentRemoval(t *testing.T) {
}

func Test_Update_HandlesRemovedComponentsFirstBeforeAddingNew(t *testing.T) {
}
