package behaviors

import (
	"components"
	"core"
	"input"
)

type Input struct {
	inputQueue input.InputQueue
	entitySet  *core.EntitySet
}

func NewInput(queue input.InputQueue, entityDB *core.EntityDB) *Input {
	input := new(Input)
	input.inputQueue = queue
	input.entitySet = entityDB.RegisterListener(input, components.INPUT)

	return input
}

// EntityListener
func (self *Input) SetUpEntity(entity *core.Entity) {
}

func (self *Input) Update(deltaT float32) {
	var input *components.Input
	nextEvents := self.inputQueue.RecentEvents()

	for _, entity := range self.entitySet.Entities {
		input = components.GetInput(entity)

		for _, event := range nextEvents {
			if callback, ok := input.Mapping[event.EventType]; ok {
				callback(entity, event)
			}
		}
	}
}
