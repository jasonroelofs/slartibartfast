package behaviors

import (
	"components"
	"core"
	"events"
	"github.com/stretchrcom/testify/assert"
	"input"
	"testing"
)

type TestInputQueue struct {
	Events input.EventList
}

func (self *TestInputQueue) RecentEvents() input.EventList {
	return self.Events
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
