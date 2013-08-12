package core

import (
	"components"
)

type EntityDatabase interface {
	ComponentAdded(*Entity)
	ComponentRemoved(*Entity, components.Component)
	EntityDestroyed(*Entity)
}

// EntityDB is as the name implies the database of entities in the system.
// Entities can have any number of Components on them, and Behaviors interact
// with Entities depending on their Components. EntityDB handles the routing
// and management logic behind all this.
//
// Logic is implemented on Entities through Listeners. Listeners must implement
// the EntityListener interface. Register a new Listener via RegisterListener, then give
// the Listener the resulting EntitySet. This EntitySet will be managed by the EntityDB and
// will always contain exactly and only the set of Entities the Listener can work with
// (according to the Component Map defined when calling RegisterListener).
//
// See the behaviors package for currently implemented Listeners.
type EntityDB struct {
	// Implements EntityDatabase

	allEntities *EntitySet
	listeners   []listenerRecord

	nextEntityId int
}

// Struct to keep track of an entity listener, the component map it can work with,
// and the set of Entities this listener should know about for processing.
type listenerRecord struct {
	listener     EntityListener
	componentMap components.ComponentType
	entitySet    *EntitySet
}

// Interface all Entity Listeners must adhere to
type EntityListener interface {
	SetUpEntity(entity *Entity)
	TearDownEntity(entity *Entity)
}

// NewEntityDB returns a new, empty EntityDB
func NewEntityDB() *EntityDB {
	return &EntityDB{
		allEntities:  NewEntitySet(),
		nextEntityId: 1,
	}
}

// RegisterListener registers the given listener to receive events and notifications
// when entities are processed through the system that contain the given set of components
func (self *EntityDB) RegisterListener(
	listener EntityListener, componentTypes ...components.ComponentType) *EntitySet {
	record := listenerRecord{listener: listener}
	record.entitySet = NewEntitySet()

	for _, ct := range componentTypes {
		record.componentMap |= ct
	}

	self.listeners = append(self.listeners, record)
	return record.entitySet
}

// RegisterEntity saves and processes a given Entity, including it into the system.
func (self *EntityDB) RegisterEntity(entity *Entity) {
	entity.Id = self.nextEntityId
	entity.entityDB = self
	self.nextEntityId++

	self.allEntities.Append(entity)
	self.notifyListenersOfNewEntity(entity)
}

// ComponentAdded is called by Entities when they recieve a new Component
// Tells all pertinent listeners that they have a new Entity to process
func (self *EntityDB) ComponentAdded(entity *Entity) {
	self.notifyListenersOfNewEntity(entity)
}

// ComponentRemoved is called by Entities when they lose a Component
// Tells all pertinent listeners that they need to clean up and ignore this Entity
func (self *EntityDB) ComponentRemoved(entity *Entity, removedComponent components.Component) {
	self.notifyListenersOfRemovedComponent(entity, removedComponent)
}

// EntityDestroyed is called by Entities when they are destroyed
// Tells all pertinent listeners to clean up and ignore this Entity
func (self *EntityDB) EntityDestroyed(entity *Entity) {
	self.notifyListenersOfDestroyedEntity(entity)
}

func (self *EntityDB) notifyListenersOfNewEntity(entity *Entity) {
	for _, listenerEntry := range self.listeners {
		if self.listenerWantsEntity(listenerEntry, entity) {
			// We don't want to re-set-up entities the Listener may already know about
			// This will happen when adding / removing components from an existing entity.
			if listenerEntry.entitySet.Append(entity) {
				listenerEntry.listener.SetUpEntity(entity)
			}
		}
	}
}

func (self *EntityDB) notifyListenersOfDestroyedEntity(entity *Entity) {
	for _, listenerEntry := range self.listeners {
		if listenerEntry.entitySet.Contains(entity) {
			listenerEntry.entitySet.Delete(entity)
			listenerEntry.listener.TearDownEntity(entity)
		}
	}
}

func (self *EntityDB) notifyListenersOfRemovedComponent(entity *Entity, component components.Component) {
	for _, listenerEntry := range self.listeners {
		// Find listeners who know about this Entity but no longer want this
		// entity (because the entity's component map no longer matches the Listener's map)
		if listenerEntry.entitySet.Contains(entity) &&
			(listenerEntry.componentMap & component.Type()) > 0 {

			listenerEntry.entitySet.Delete(entity)
			listenerEntry.listener.TearDownEntity(entity)
		}
	}
}

func (self *EntityDB) listenerWantsEntity(le listenerRecord, entity *Entity) bool {
	return (entity.ComponentMap() & le.componentMap) == le.componentMap
}
