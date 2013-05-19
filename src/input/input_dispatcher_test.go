package input

import (
	"events"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewInputDispatcher(t *testing.T) {
	mapper := NewInputDispatcher()
	assert.NotNil(t, mapper, "Mapper failed to initialize")
}

func Test_On_RegistersACallbackForAnEvent(t *testing.T) {
	mapper := NewInputDispatcher()
	quitCalled := false

	mapper.On(events.Quit, func(events.Event) {
		quitCalled = true
	})

	mapper.keyCallback(KeyQ, 1)

	assert.True(t, quitCalled, "Quit callback was not called")
}

func Test_DoesNothingIfNoEventForKey(t *testing.T) {
	mapper := NewInputDispatcher()

	assert.NotPanics(t, func() {
		mapper.keyCallback(KeyQ, 1)
	}, "Mapper panic'd on unmapped key")
}

func Test_DoesNothingIfNoKeyMappedToEvent(t *testing.T) {
	mapper := NewInputDispatcher()
	jumpEvent := events.EventType(100)
	jumpCalled := false

	mapper.On(jumpEvent, func(events.Event) {
		jumpCalled = true
	})

	mapper.keyCallback(KeyN, 1)

	assert.False(t, jumpCalled, "Jump event called when it should not have")
}

func Test_CanMapMultipleKeysToOneEvent(t *testing.T) {
	mapper := NewInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyEsc, events.Quit)

	quitCallCount := 0

	mapper.On(events.Quit, func(events.Event) {
		quitCallCount += 1
	})

	mapper.keyCallback(KeyQ, 1)
	mapper.keyCallback(KeyEsc, 1)

	assert.Equal(t, 2, quitCallCount)
}

func Test_StoresKeyEventsReceived(t *testing.T) {
	mapper := NewInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	mapper.keyCallback(KeyQ, 1)
	mapper.keyCallback(KeyD, 0)

	assert.Equal(t, 2, len(mapper.storedEvents))
	assert.Equal(t, events.Quit, mapper.storedEvents[0].EventType)
	assert.True(t, mapper.storedEvents[0].Pressed)

	assert.Equal(t, events.MoveForward, mapper.storedEvents[1].EventType)
	assert.False(t, mapper.storedEvents[1].Pressed)
}

func Test_EventsSinceLast_ReturnsStoredEventsAndClearsList(t *testing.T) {
	mapper := NewInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	mapper.keyCallback(KeyQ, 1)
	mapper.keyCallback(KeyD, 0)

	nextEvents := mapper.RecentEvents()

	assert.Equal(t, 2, len(nextEvents))
	assert.Equal(t, 0, len(mapper.storedEvents))
}
