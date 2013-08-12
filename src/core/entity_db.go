package core

import (
	"components"
)

type EntityDatabase interface {
	ComponentAdded(*Entity)
	ComponentRemoved(*Entity, components.Component)
	EntityDestroyed(*Entity)
}

// EntityDB is as the name implies the database of entities in the system
// Entities can have any number of Components on them, and Behaviors interact
// with Entities depending on their Components. EntityDB handles the routing
// and management logic behind all this
type EntityDB struct {
	// Implements EntityDatabase

	allEntities *EntitySet
	listeners   []listenerRecord

	nextEntityId int
}

// Struct to keep track of a listener and the set of
// components said Listener registered to receive notifications of
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

// NewEntityDB returns a new EntityDB
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

// RegisterEntity saves and processes a given Entity for inclusion in the system.
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
			listenerEntry.entitySet.Append(entity)
			listenerEntry.listener.SetUpEntity(entity)
		}
	}
}

func (self *EntityDB) notifyListenersOfDestroyedEntity(entity *Entity) {
	for _, listenerEntry := range self.listeners {
		if self.listenerWantsEntity(listenerEntry, entity) {
			listenerEntry.entitySet.Delete(entity)
			listenerEntry.listener.TearDownEntity(entity)
		}
	}
}

func (self *EntityDB) notifyListenersOfRemovedComponent(entity *Entity, component components.Component) {
	for _, listenerEntry := range self.listeners {
		// Find listeners who know about this Entity but no longer want this
		// entity (because the component map no longer matches the requested map)
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
