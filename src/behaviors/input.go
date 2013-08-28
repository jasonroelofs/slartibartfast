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

// SetUpEntity :: EntityListener
// For input components that want event polling, tells the InputQueue to
// create Events using Polling instead of just callbacks
func (self *Input) SetUpEntity(entity *core.Entity) {
	input := components.GetInput(entity)

	if input.WantsPolling() {
		self.inputQueue.PollEvents(input.Polling)
	}
}

// TearDownEntity :: EntityListener
func (self *Input) TearDownEntity(entity *core.Entity) {
	input := components.GetInput(entity)

	if input.WantsPolling() {
		self.inputQueue.UnpollEvents(input.Polling)
	}
}

func (self *Input) Update(deltaT float32) {
	var input *components.Input
	nextEvents := self.inputQueue.RecentEvents()

	for _, entity := range self.entitySet.Entities() {
		input = components.GetInput(entity)

		for _, event := range nextEvents {
			if callback, ok := input.Mapping[event.EventType]; ok {
				callback(entity, event)
			}
		}
	}
}
