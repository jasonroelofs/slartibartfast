package core

import (
	"components"
)

// EntityDB is as the name implies the database of entities in the system
// Entities can have any number of Components on them, and Behaviors interact
// with Entities depending on their Components. EntityDB handles the routing
// and management logic behind all this
type EntityDB struct {
	allEntities EntitySet
	listeners   []listenerRecord
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

// RegisterEntity saves and processes a given Entity for inclusion in the system.
func (self *EntityDB) RegisterEntity(entity *Entity) {
	self.allEntities.Append(entity)
	self.notifyListenersOfNewEntity(entity)

	// Clear entity flags relating added and removed components
	// so Update() doesn't re-apply the initial state of the Entity
	entity.finalizeComponentAddition()
	entity.finalizeComponentRemoval()
}

// RegisterListener registers the given listener to receive events and notifications
// when entities are processed through the system that contain the given set of components
func (self *EntityDB) RegisterListener(
	listener EntityListener, componentTypes ...components.ComponentType) *EntitySet {
	record := listenerRecord{listener: listener}
	record.entitySet = new(EntitySet)

	for _, ct := range componentTypes {
		record.componentMap |= ct
	}

	self.listeners = append(self.listeners, record)
	return record.entitySet
}

// Update is called every frame and checks for dirty and/or delete-flagged Entities
func (self *EntityDB) Update() {
	for _, entity := range self.allEntities.Entities() {
		if entity.destroyNextFrame {
			self.notifyListenersOfDeletedEntity(entity)
			self.allEntities.Delete(entity)
		}

		if entity.anyComponentsRemoved() {
			self.notifyListenersOfChangedComponents(entity)
			entity.finalizeComponentRemoval()
		}

		if entity.anyComponentsAdded() {
			self.notifyListenersOfNewEntity(entity)
			entity.finalizeComponentAddition()
		}
	}
}

func (self *EntityDB) notifyListenersOfNewEntity(entity *Entity) {
	for _, listenerEntry := range self.listeners {
		if self.listenerWantsEntity(listenerEntry, entity) {
			listenerEntry.entitySet.Append(entity)
			listenerEntry.listener.SetUpEntity(entity)
		}
	}
}

func (self *EntityDB) notifyListenersOfDeletedEntity(entity *Entity) {
	for _, listenerEntry := range self.listeners {
		// Find only the listeners who manage components this Entity uses
		if self.listenerWantsEntity(listenerEntry, entity) {
			listenerEntry.entitySet.Delete(entity)
			listenerEntry.listener.TearDownEntity(entity)
		}
	}
}

func (self *EntityDB) notifyListenersOfChangedComponents(entity *Entity) {
	for _, listenerEntry := range self.listeners {
		// Find listeners who know about this Entity but no longer want this
		// entity (because the component map no longer matches the requested map)
		if listenerEntry.entitySet.Contains(entity) &&
			!self.listenerWantsEntity(listenerEntry, entity) {

			listenerEntry.entitySet.Delete(entity)
			listenerEntry.listener.TearDownEntity(entity)
		}
	}
}

func (self *EntityDB) listenerWantsEntity(le listenerRecord, entity *Entity) bool {
	return (entity.ComponentMap() & le.componentMap) == le.componentMap
}
