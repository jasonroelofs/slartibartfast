package core

// A set of entities, most often to be a subset of all entities
// in the system. EntityDB.RegisterListener returns this object and will
// keep the list of Entities up to date with Entity changes
type EntitySet struct {
	entities []*Entity
}

// Append adds the given Entity to the set
func (self *EntitySet) Append(entity *Entity) {
	self.entities = append(self.entities, entity)
}

// Len returns the number of entities in this set
func (self *EntitySet) Len() int {
	return len(self.entities)
}

// Get returns the entity at the given Index.
// Does not have any protections against out-of-bounds errors right now.
func (self *EntitySet) Get(index int) *Entity {
	return self.entities[index]
}

// Entities returns the array of entities in this set, mainly for iterating
// over. Might turn this into more of a callback loop rather than giving
// the calling code a raw array.
func (self *EntitySet) Entities() []*Entity {
	return self.entities
}
