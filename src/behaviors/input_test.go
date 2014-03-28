package behaviors

import (
	"components"
	"core"
	"events"
	"github.com/stretchr/testify/assert"
	"input"
	"testing"
)

type TestInputQueue struct {
	Events          input.EventList
	pollingEvents   []events.EventType
	unpollingEvents []events.EventType
}

func (self *TestInputQueue) RecentEvents() input.EventList {
	return self.Events
}

func (self *TestInputQueue) PollEvents(events []events.EventType) {
	self.pollingEvents = events
}

func (self *TestInputQueue) UnpollEvents(events []events.EventType) {
	self.unpollingEvents = events
}

func getTestInput() (*Input, *TestInputQueue, *core.EntityDB) {
	entityDB := core.NewEntityDB()
	queue := new(TestInputQueue)

	input := NewInput(queue, entityDB)

	return input, queue, entityDB
}

func Test_NewInput(t *testing.T) {
	input, queue, _ := getTestInput()

	assert.NotNil(t, input.entitySet)
	assert.Equal(t, queue, input.inputQueue)
}

func Test_SetUpEntity_TellsQueueWhatEventsToPollFor(t *testing.T) {
	_, queue, entityDb := getTestInput()

	pollEvents := []events.EventType{events.Quit, events.TurnLeft}
	entity := core.NewEntity()
	entity.AddComponent(&components.Input{
		Polling: pollEvents,
	})
	entityDb.RegisterEntity(entity)

	assert.Equal(t, pollEvents, queue.pollingEvents)
}

func Test_TearDownEntity_TurnsOffPollingForRelatedEvents(t *testing.T) {
	_, queue, entityDb := getTestInput()

	pollEvents := []events.EventType{events.Quit, events.TurnLeft}
	entity := core.NewEntity()
	entity.AddComponent(&components.Input{
		Polling: pollEvents,
	})
	entityDb.RegisterEntity(entity)
	entityDb.EntityDestroyed(entity)

	assert.Equal(t, pollEvents, queue.unpollingEvents)
}

func Test_Update_PassesInputEventsToComponentsWhoWantThem(t *testing.T) {
	input, queue, entityDb := getTestInput()

	testMappingQuitCalled := false
	var testMappingQuitCalledEntity components.ComponentHolder

	testMapping := components.InputEventMap{
		events.Quit: func(entity components.ComponentHolder, event events.Event) {
			testMappingQuitCalledEntity = entity
			testMappingQuitCalled = true
		},
	}

	entity := core.NewEntity()
	entity.AddComponent(&components.Input{
		Mapping: testMapping,
	})
	entityDb.RegisterEntity(entity)

	queue.Events = append(queue.Events, events.Event{EventType: events.Quit})

	input.Update(0)

	assert.True(t, testMappingQuitCalled)
	assert.Equal(t, entity, testMappingQuitCalledEntity)
}
