package core

// A set of entities, most often to be a subset of all entities
// in the system. EntityDB.RegisterListener returns this object and will
// keep the list of Entities up to date with Entity changes
type EntitySet struct {
	entities map[int]*Entity
}

// NewEntitySet initializes and returns an empty EntitySet
func NewEntitySet() *EntitySet {
	return &EntitySet{
		entities: make(map[int]*Entity),
	}
}

// Append adds the given Entity to the set
// Does not allow duplicate Entity records in the set.
// Returns whether the Entity was added to the set or not.
func (self *EntitySet) Append(entity *Entity) (appended bool) {
	if _, ok := self.entities[entity.Id]; !ok {
		self.entities[entity.Id] = entity
		appended = true
	} else {
		appended = false
	}

	return
}

// Len returns the number of entities in this set
func (self *EntitySet) Len() int {
	return len(self.entities)
}

// Get returns the entity of the given Id.
// If no Entity of that Id is in this Set, returns nil
func (self *EntitySet) Get(entityId int) *Entity {
	return self.entities[entityId]
}

// Entities returns the array of entities in this set, mainly for iterating
// over. Might turn this into more of a callback loop rather than giving
// the calling code a raw array.
// TODO Currently inefficient, creates a new array every time it's called
func (self *EntitySet) Entities() (entityList []*Entity) {
	for _, entity := range self.entities {
		entityList = append(entityList, entity)
	}
	return
}

// Delete removes the given Entity from this set
func (self *EntitySet) Delete(entity *Entity) {
	delete(self.entities, entity.Id)
}

// Contains returns true or false depending on if the given Entity
// is present in this set
func (self *EntitySet) Contains(entity *Entity) bool {
	_, ok := self.entities[entity.Id]
	return ok
}
