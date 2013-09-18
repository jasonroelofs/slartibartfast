package input

import (
	"events"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func init() {
	// To not confuse tests with the current set of hard-coded key - event mappings
	InputDispatcherTesting = true
}

type TestEmitter struct {
	// Implements input.InputEmitter

	keyCallback      func(KeyCode, KeyState)
	mousePosCallback func(int, int)

	keyStates map[KeyCode]KeyState
}

func (self *TestEmitter) KeyCallback(cb func(KeyCode, KeyState)) {
	self.keyCallback = cb
}

func (self *TestEmitter) fireKeyCallback(key KeyCode, state KeyState) {
	self.keyCallback(key, state)
}

func (self *TestEmitter) MouseButtonCallback(cb func(int, KeyState)) {
}

func (self *TestEmitter) MousePositionCallback(cb func(int, int)) {
	self.mousePosCallback = cb
}

func (self *TestEmitter) moveMouse(x, y int) {
	self.mousePosCallback(x, y)
}

func (self *TestEmitter) MouseWheelCallback(cb func(int)) {
}

func (self *TestEmitter) IsKeyPressed(key KeyCode) bool {
	return self.keyStates[key] == KeyPressed
}

func (self *TestEmitter) setKeyState(key KeyCode, state KeyState) {
	self.keyStates[key] = state
}

func NewTestEmitter() *TestEmitter {
	return &TestEmitter{
		keyStates: make(map[KeyCode]KeyState),
	}
}

func GetInputDispatcher() (*InputDispatcher, *TestEmitter) {
	config, _ := configs.NewConfig("testdata/blank.json")
	emitter := NewTestEmitter()
	return NewInputDispatcher(config, emitter), emitter
}

func Test_NewInputDispatcher(t *testing.T) {
	mapper, _ := GetInputDispatcher()
	assert.NotNil(t, mapper, "Mapper failed to initialize")
}

func Test_On_RegistersACallbackForEventFiring(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	quitCalled := false
	mapper.mapKeyToEvent(KeyQ, events.Quit)

	mapper.On(events.Quit, func(events.Event) {
		quitCalled = true
	})

	emitter.fireKeyCallback(KeyQ, KeyPressed)

	assert.True(t, quitCalled, "Quit callback was not called")
}

func Test_OnKey_RegistersCallbackForRawKeyEvent(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	pKeyHit := false
	var pKeyEvent events.Event

	mapper.OnKey(KeyP, func(event events.Event) {
		pKeyHit = true
		pKeyEvent = event
	})

	emitter.fireKeyCallback(KeyP, KeyPressed)

	assert.True(t, pKeyHit, "P key callback was not called")
	assert.True(t, pKeyEvent.Pressed)
}

func Test_OnKey_CanMapMultipleCallbacksForASingleKey(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	callback1Hit := false
	callback2Hit := false

	mapper.OnKey(KeyP, func(event events.Event) {
		callback1Hit = true
	})

	mapper.OnKey(KeyP, func(event events.Event) {
		callback2Hit = true
	})

	emitter.fireKeyCallback(KeyP, KeyPressed)

	assert.True(t, callback1Hit, "Did not call the first callback")
	assert.True(t, callback2Hit, "Did not call the second callback")
}

func Test_IgnoresUnmappedKeyEvents(t *testing.T) {
	_, emitter := GetInputDispatcher()
	emitter.fireKeyCallback(KeyQ, KeyPressed)
}

func Test_DoesNothingIfNoKeyMappedToMappedEventCallback(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	jumpEvent := events.EventType(100)
	jumpCalled := false

	mapper.On(jumpEvent, func(events.Event) {
		jumpCalled = true
	})

	emitter.fireKeyCallback(KeyN, KeyPressed)

	assert.False(t, jumpCalled, "Jump event called when it should not have")
}

func Test_On_CanMapMultipleKeysToOneEvent(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyEsc, events.Quit)

	quitCallCount := 0

	mapper.On(events.Quit, func(events.Event) {
		quitCallCount += 1
	})

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyEsc, KeyPressed)

	assert.Equal(t, 2, quitCallCount)
}

func Test_On_CanMapMultipleEventsToOneKey(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
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

	emitter.fireKeyCallback(KeyJ, KeyPressed)

	assert.Equal(t, 1, forwardCallCount)
	assert.Equal(t, 1, backwardCallCount)
}

func Test_RecentEvents_ReturnsLastSetOfReceivedKeyEvents(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyD, KeyReleased)

	recentEvents := mapper.RecentEvents()

	assert.Equal(t, 2, len(recentEvents))
	assert.Equal(t, events.Quit, recentEvents[0].EventType)
	assert.True(t, recentEvents[0].Pressed)

	assert.Equal(t, events.MoveForward, recentEvents[1].EventType)
	assert.False(t, recentEvents[1].Pressed)
}

func Test_RecentEvents_ReturnsLastSetOfReceivedMouseEvents(t *testing.T) {
	mapper, emitter := GetInputDispatcher()

	// Mouse coords are transformed from 0,0 top left to 0,0 center
	emitter.moveMouse(10, 20)
	emitter.moveMouse(-5, 13)

	recentEvents := mapper.RecentEvents()

	assert.Equal(t, 2, len(recentEvents))
	assert.Equal(t, events.MouseMove, recentEvents[0].EventType)
	assert.Equal(t, 10, recentEvents[0].MouseXDiff)
	assert.Equal(t, 20, recentEvents[0].MouseYDiff)

	assert.Equal(t, events.MouseMove, recentEvents[1].EventType)
	assert.Equal(t, -5, recentEvents[1].MouseXDiff)
	assert.Equal(t, 13, recentEvents[1].MouseYDiff)
}

func Test_RecentEvents_ClearsListForNextCall(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyD, KeyReleased)

	nextEvents1 := mapper.RecentEvents()
	nextEvents2 := mapper.RecentEvents()

	assert.Equal(t, 2, len(nextEvents1))
	assert.Equal(t, 0, len(nextEvents2))
}

func Test_PollEvents_IncludesEventsInRecentEventsIfKeyPressed(t *testing.T) {
}

// TODO: Merge these tests into how RecentEvents works so it's testing
// behavior and not implementation.
func Test_PollEvents_AddsToListOfEventsToPoll(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
	}
	mapper.PollEvents(eventList)

	emitter.setKeyState(KeyQ, KeyPressed)
	emitter.setKeyState(KeyD, KeyPressed)

	nextEvents := mapper.RecentEvents()

	assert.Equal(t, 2, len(nextEvents))
	assert.Equal(t, events.Quit, nextEvents[0].EventType)
	assert.Equal(t, events.MoveForward, nextEvents[1].EventType)
}

func Test_UnpollEvents_RemovesKnownPollEventsFromPolling(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
		events.MoveBackward,
	}
	mapper.PollEvents(eventList)
	mapper.UnpollEvents([]events.EventType{events.Quit})

	emitter.setKeyState(KeyQ, KeyPressed)
	emitter.setKeyState(KeyD, KeyPressed)

	nextEvents := mapper.RecentEvents()

	assert.Equal(t, 1, len(nextEvents))
	assert.Equal(t, events.MoveForward, nextEvents[0].EventType)
}
