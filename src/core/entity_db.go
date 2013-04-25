package core

import (
	"components"
)

// EntityDB is as the name implies the database of entities in the system
// Entities can have any number of Components on them, and Behaviors interact
// with Entities depending on their Components. EntityDB handles the routing
// and management logic behind all this

type EntityListener interface {
	SetUpEntity(entity *Entity)
}

type EntityList []Entity

// Struct to keep track of a listener and the set of
// components said Listener registered to receive notifications of
type listenerRecord struct {
	listener     EntityListener
	componentMap components.ComponentType
	entities     EntityList
}

type EntityDB struct {
	allEntities []Entity
	listeners   []listenerRecord
}

// RegisterEntity saves and processes a given Entity for inclusion in the system.
func (self *EntityDB) RegisterEntity(entity *Entity) {
	self.allEntities = append(self.allEntities, *entity)
	self.notifyListenersOfNewEntity(entity)
}

// RegisterListener registers the given listener to receive events and notifications
// when entities are processed through the system that contain the given set of components
func (self *EntityDB) RegisterListener(
	listener EntityListener, componentTypes ...components.ComponentType) EntityList {
	record := listenerRecord{listener: listener}

	for _, ct := range componentTypes {
		record.componentMap |= ct
	}

	self.listeners = append(self.listeners, record)
	return record.entities
}

func (self *EntityDB) notifyListenersOfNewEntity(entity *Entity) {
	for _, entry := range self.listeners {
		if entity.ComponentMap() & entry.componentMap > 0 {
			entry.listener.SetUpEntity(entity)
		}
	}
}
