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
func Test_RegisterEntity_KeepsTrackOfEntitiesAndComponents(t *testing.T) {
	db := EntityDB{}
	entity := NewEntity()
	entity.AddComponent(new(components.Visual))

	db.RegisterEntity(entity)

	assert.Equal(t, 1, db.allEntities.Len())
}

func Test_RegisterEntity_GivesSelfPointerToEntity(t *testing.T) {
	db := EntityDB{}
	entity := NewEntity()
	db.RegisterEntity(entity)

	assert.Equal(t, &db, entity.entityDB)
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
func Test_RegisterEntity_ListenerNotifiedOfNewEntityMatchingComponents(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entity := NewEntity()

	db.RegisterListener(listener, components.TRANSFORM)
	db.RegisterEntity(entity)

	assert.Equal(t, 1, len(listener.setUpEntities))
	assert.Equal(t, entity, listener.setUpEntities[0])
}

func Test_RegisterEntity_ListenerNotNotifiedOfNewEntityIfComponentsDontMatch(t *testing.T) {
	db := EntityDB{}
	listener := new(TestListener)
	entity := Entity{}

	db.RegisterListener(listener, components.TRANSFORM)
	db.RegisterEntity(&entity)

	assert.Equal(t, 0, len(listener.setUpEntities))
}

func Test_RegisterEntity_AddsEntityToListenerEntityListIfComponentsMatch(t *testing.T) {
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
// Entity Processing (EntityDatabase interface)
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

func Test_EntityDestroyed_RemovesEntityFromEntitySets(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	db.RegisterEntity(entity)

	entity.Destroy()

	assert.Equal(t, 0, listeners[0].entitySet.Len())
	assert.Equal(t, 0, listeners[1].entitySet.Len())
	assert.Equal(t, 0, listeners[2].entitySet.Len())
}

func Test_EntityDestroyed_TellsListenersToCleanUpEntities(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	db.RegisterEntity(entity)

	entity.Destroy()

	// Triggers on listeners who want the entity
	assert.Equal(t, entity, listeners[0].tearDownEntities[0])
	// Ignores all others
	assert.Equal(t, 0, len(listeners[1].tearDownEntities))
	assert.Equal(t, 0, len(listeners[2].tearDownEntities))
}

func Test_ComponentAdded_TellsListenersToSetUpNewEntityOnNewComponents(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	db.RegisterEntity(entity)

	entity.AddComponent(new(components.Visual))

	// Check set up callback triggered
	assert.Equal(t, entity, listeners[1].setUpEntities[0])
	// And added to the listener's entity set
	assert.Equal(t, entity, listeners[1].entitySet.Entities()[0])
}

func Test_ComponentRemoved_TellsListenersToTearDownEntityOnComponentRemoval(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	entity.AddComponent(new(components.Visual))
	db.RegisterEntity(entity)

	entity.RemoveComponent(components.VISUAL)

	// Check tear down callback triggered
	assert.Equal(t, entity, listeners[1].tearDownEntities[0])
	// And removed from the listeners' entity set
	assert.Equal(t, 0, listeners[1].entitySet.Len())

	// Sanity, no other listener got the callback
	assert.Equal(t, 0, len(listeners[0].tearDownEntities))
	assert.Equal(t, 0, len(listeners[2].tearDownEntities))
	assert.Equal(t, 1, listeners[0].entitySet.Len())
	assert.Equal(t, 0, listeners[2].entitySet.Len())
}

func Test_Update_HandlesComponentReplacementProperly(t *testing.T) {
	db, listeners := RegisterDBAndListeners()
	entity := NewEntity()
	oldVisual := new(components.Visual)
	entity.AddComponent(oldVisual)
	db.RegisterEntity(entity)

	// Now replace the oldVisual with a new visual component
	entity.RemoveComponent(components.VISUAL)
	newVisual := new(components.Visual)
	entity.AddComponent(newVisual)

	assert.Equal(t, entity, listeners[1].tearDownEntities[0])
	assert.Equal(t, entity, listeners[1].setUpEntities[0])

	// Entity still stored
	assert.True(t, listeners[1].entitySet.Contains(entity))
	// Using the new visual component (move test elsewhere?)
	assert.Equal(t, newVisual, components.GetVisual(entity))
}
