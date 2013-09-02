package input

import (
	"events"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

// TODO
// Rewrite a bunch of this to focus on behavior now that input emission is
// abstracted into an interface

func init() {
	InputDispatcherTesting = true
}

type TestEmitter struct {
	// Implements input.InputEmitter
}

func (self TestEmitter) KeyCallback(cb func(int, KeyState)) {
}

func (self TestEmitter) MouseButtonCallback(cb func(int, KeyState)) {
}

func (self TestEmitter) MousePositionCallback(cb func(int, int)) {
}

func (self TestEmitter) MouseWheelCallback(cb func(int)) {
}

func (self TestEmitter) IsKeyPressed(key int) bool {
	return false
}

func GetInputDispatcher() *InputDispatcher {
	emitter := new(TestEmitter)
	return NewInputDispatcher(emitter)
}

func Test_NewInputDispatcher(t *testing.T) {
	mapper := GetInputDispatcher()
	assert.NotNil(t, mapper, "Mapper failed to initialize")
}

func Test_On_RegistersACallbackForAnEvent(t *testing.T) {
	mapper := GetInputDispatcher()
	quitCalled := false
	mapper.mapKeyToEvent(KeyQ, events.Quit)

	mapper.On(events.Quit, func(events.Event) {
		quitCalled = true
	})

	mapper.keyCallback(KeyQ, 1)

	assert.True(t, quitCalled, "Quit callback was not called")
}

func Test_OnKey_RegistersCallbackForRawKeyEvent(t *testing.T) {
	mapper := GetInputDispatcher()
	pKeyHit := false
	var pKeyEvent events.Event

	mapper.OnKey(KeyP, func(event events.Event) {
		pKeyHit = true
		pKeyEvent = event
	})

	mapper.keyCallback(KeyP, 1)

	assert.True(t, pKeyHit, "P key callback was not called")
	assert.True(t, pKeyEvent.Pressed)
}

func Test_OnKey_CanMapMultipleCallbacksForASingleKey(t *testing.T) {
	mapper := GetInputDispatcher()
	callback1Hit := false
	callback2Hit := false

	mapper.OnKey(KeyP, func(event events.Event) {
		callback1Hit = true
	})

	mapper.OnKey(KeyP, func(event events.Event) {
		callback2Hit = true
	})

	mapper.keyCallback(KeyP, 1)

	assert.True(t, callback1Hit, "Did not call the first callback")
	assert.True(t, callback2Hit, "Did not call the second callback")
}

func Test_DoesNothingIfNoEventForKey(t *testing.T) {
	mapper := GetInputDispatcher()

	assert.NotPanics(t, func() {
		mapper.keyCallback(KeyQ, 1)
	}, "Mapper panic'd on unmapped key")
}

func Test_DoesNothingIfNoKeyMappedToEvent(t *testing.T) {
	mapper := GetInputDispatcher()
	jumpEvent := events.EventType(100)
	jumpCalled := false

	mapper.On(jumpEvent, func(events.Event) {
		jumpCalled = true
	})

	mapper.keyCallback(KeyN, 1)

	assert.False(t, jumpCalled, "Jump event called when it should not have")
}

func Test_On_CanMapMultipleKeysToOneEvent(t *testing.T) {
	mapper := GetInputDispatcher()
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

func Test_On_CanMapMultipleEventsToOneKey(t *testing.T) {
	mapper := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyJ, events.MoveForward)
	mapper.mapKeyToEvent(KeyJ, events.MoveBackward)

	forwardCallCount := 0
	backwardCallCount := 0

	mapper.On(events.MoveForward, func(events.Event) {
		forwardCallCount += 1
	})

	mapper.On(events.MoveBackward, func(events.Event) {
		backwardCallCount += 1
	})

	mapper.keyCallback(KeyJ, 1)

	assert.Equal(t, 1, forwardCallCount)
	assert.Equal(t, 1, backwardCallCount)
}

func Test_StoresKeyEventsReceived(t *testing.T) {
	mapper := GetInputDispatcher()
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

func Test_StoresMouseEventsReceived(t *testing.T) {
	mapper := GetInputDispatcher()

	// Mouse coords are transformed from 0,0 top left to 0,0 center
	mapper.mouseMoveCallback(10, 20)
	mapper.mouseMoveCallback(-5, 13)

	assert.Equal(t, 2, len(mapper.storedEvents))
	assert.Equal(t, events.MouseMove, mapper.storedEvents[0].EventType)
	assert.Equal(t, 10, mapper.storedEvents[0].MouseXDiff)
	assert.Equal(t, 20, mapper.storedEvents[0].MouseYDiff)

	assert.Equal(t, events.MouseMove, mapper.storedEvents[1].EventType)
	assert.Equal(t, -5, mapper.storedEvents[1].MouseXDiff)
	assert.Equal(t, 13, mapper.storedEvents[1].MouseYDiff)
}

func Test_RecentEvents_ReturnsStoredEventsAndClearsList(t *testing.T) {
	mapper := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	mapper.keyCallback(KeyQ, 1)
	mapper.keyCallback(KeyD, 0)

	nextEvents := mapper.RecentEvents()

	assert.Equal(t, 2, len(nextEvents))
	assert.Equal(t, 0, len(mapper.storedEvents))
}

func Test_RecentEvents_ReturnsEmptyListOnNoEvents(t *testing.T) {
	mapper := GetInputDispatcher()
	nextEvents := mapper.RecentEvents()

	assert.Equal(t, 0, len(nextEvents))
}

// TODO: Merge these tests into how RecentEvents works so it's testing
// behavior and not implementation.
func Test_PollEvents_AddsToListOfEventsToPoll(t *testing.T) {
	mapper := GetInputDispatcher()
	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
		events.MoveBackward,
	}
	mapper.PollEvents(eventList)

	assert.True(t, mapper.pollingEvents[events.Quit])
	assert.True(t, mapper.pollingEvents[events.MoveForward])
	assert.True(t, mapper.pollingEvents[events.MoveBackward])
}

func Test_UnpollEvents(t *testing.T) {
	mapper := GetInputDispatcher()
	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
		events.MoveBackward,
	}
	mapper.PollEvents(eventList)

	mapper.UnpollEvents([]events.EventType{events.Quit})

	assert.False(t, mapper.pollingEvents[events.Quit])
	assert.True(t, mapper.pollingEvents[events.MoveForward])
	assert.True(t, mapper.pollingEvents[events.MoveBackward])
}

// Not sure how to test this without abstracting glfw
//func Test_RecentEvents_IncludesEventsPollingSaysAreFiring(t *testing.T) {
//	mapper := GetInputDispatcher()
//	eventList := []events.EventType{
//		events.MoveForward,
//	}
//	mapper.PollEvents(eventList)
//
//	mapper.mapKeyToEvent(KeyQ, events.Quit)
//	mapper.mapKeyToEvent(KeyD, events.MoveForward)
//
//	nextEvents := mapper.RecentEvents()
//
//	assert.Equal(t, 2, len(nextEvents))
//}
