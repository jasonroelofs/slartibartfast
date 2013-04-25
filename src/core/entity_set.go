package core

// A set of entities, most often to be a subset of all entities
// in the system. EntityDB.RegisterListener returns this object and will
// keep the list of Entities up to date with Entity changes
type EntitySet struct {
	Entities []*Entity
}

func (self *EntitySet) Append(entity *Entity) {
	self.Entities = append(self.Entities, entity)
}
