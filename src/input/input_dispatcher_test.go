package input

import (
	"configs"
	"events"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestEmitter struct {
	// Implements input.InputEmitter

	keyCallback      func(KeyCode, KeyState)
	mousePosCallback func(int, int)

	mouseButtonCallback func(MouseButtonCode, KeyState)

	HidingCursor bool

	keyStates map[KeyCode]KeyState
}

func (self *TestEmitter) KeyCallback(cb func(KeyCode, KeyState)) {
	self.keyCallback = cb
}

func (self *TestEmitter) fireKeyCallback(key KeyCode, state KeyState) {
	self.keyCallback(key, state)
}

func (self *TestEmitter) MouseButtonCallback(cb func(MouseButtonCode, KeyState)) {
	self.mouseButtonCallback = cb
}

func (self *TestEmitter) fireMouseButtonCallback(button MouseButtonCode, state KeyState) {
	self.mouseButtonCallback(button, state)
}

func (self *TestEmitter) MousePositionCallback(cb func(int, int)) {
	self.mousePosCallback = cb
}

func (self *TestEmitter) ShowCursor() {
	self.HidingCursor = false
}

func (self *TestEmitter) HideCursor() {
	self.HidingCursor = true
}

func (self *TestEmitter) moveMouse(x, y int) {
	self.mousePosCallback(x, y)
}

func (self *TestEmitter) MouseWheelCallback(cb func(int)) {
}

func (self *TestEmitter) IsKeyPressed(key KeyCode) bool {
	return self.keyStates[key] == KeyPressed || self.keyStates[key] == KeyRepeated
}

func (self *TestEmitter) IsKeyRepeated(key KeyCode) bool {
	return self.keyStates[key] == KeyRepeated
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

func Test_NewInputDispatcher_ReadsKeybindingsFromConfig(t *testing.T) {
	config, _ := configs.NewConfig("testdata/keybindings.json")
	emitter := NewTestEmitter()
	dispatcher := NewInputDispatcher(config, emitter)

	quitFired := 0
	dispatcher.On(events.Quit, func(events.Event) {
		quitFired += 1
	})

	zoomFired := 0
	dispatcher.On(events.ZoomIn, func(events.Event) {
		zoomFired += 1
	})

	attackFired := 0
	dispatcher.On(events.Primary, func(events.Event) {
		attackFired += 1
	})

	// See the test file keybindings.json
	// 3 Quit keys
	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyEsc, KeyPressed)
	emitter.fireKeyCallback(KeyX, KeyPressed)

	// One ZoomIn key
	emitter.fireKeyCallback(KeyZ, KeyPressed)

	// A mouse button
	emitter.fireMouseButtonCallback(Mouse1, KeyPressed)

	assert.Equal(t, 3, quitFired)
	assert.Equal(t, 1, zoomFired)
	assert.Equal(t, 1, attackFired)
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

func Test_On_MapsRepeatAndPressedAsKeyPressed(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)

	quitCallCount := 0

	mapper.On(events.Quit, func(events.Event) {
		quitCallCount += 1
	})

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyQ, KeyRepeated)

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

func Test_On_CanMapMouseButtonsToEvents(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapMouseButtonToEvent(Mouse1, events.MoveForward)
	mapper.mapMouseButtonToEvent(Mouse2, events.MoveBackward)

	forwardCallCount := 0
	backwardCallCount := 0

	mapper.On(events.MoveForward, func(events.Event) {
		forwardCallCount += 1
	})

	mapper.On(events.MoveBackward, func(events.Event) {
		backwardCallCount += 1
	})

	emitter.fireMouseButtonCallback(Mouse1, KeyPressed)
	emitter.fireMouseButtonCallback(Mouse2, KeyPressed)

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

func Test_RecentEvents_FlagsRepeatedKeysAsPressedAndRepeated(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyQ, KeyRepeated)

	recentEvents := mapper.RecentEvents()

	assert.Equal(t, 2, len(recentEvents))
	assert.True(t, recentEvents[0].Pressed)

	assert.True(t, recentEvents[1].Pressed)
	assert.True(t, recentEvents[1].Repeated)
}

func Test_RecentEvents_ReturnsLastSetOfReceivedMouseEvents(t *testing.T) {
	mapper, emitter := GetInputDispatcher()

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

func Test_HideCursor_HidesCursorAndFlagsMouseToResetToCenterAfterMove(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.HideCursor()

	assert.True(t, emitter.HidingCursor)
}

func Test_ShowCursor_ShowsCursorAndStopsResettingCursorAfterMove(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.HideCursor()
	mapper.ShowCursor()

	assert.False(t, emitter.HidingCursor)
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

func Test_PollEvents_AddsToListOfEventsToPollToRecentEvents(t *testing.T) {
	mapper, emitter := GetInputDispatcher()
	mapper.mapKeyToEvent(KeyQ, events.Quit)
	mapper.mapKeyToEvent(KeyD, events.MoveForward)
	mapper.mapKeyToEvent(KeyE, events.MoveBackward)

	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
		events.MoveBackward,
	}
	mapper.PollEvents(eventList)

	emitter.setKeyState(KeyQ, KeyPressed)
	emitter.setKeyState(KeyD, KeyPressed)
	emitter.setKeyState(KeyE, KeyRepeated)

	nextEvents := mapper.RecentEvents()

	assert.Equal(t, 3, len(nextEvents))
	assert.Equal(t, events.Quit, nextEvents[0].EventType)
	assert.Equal(t, events.MoveForward, nextEvents[1].EventType)

	// Repeat key handling
	assert.Equal(t, events.MoveBackward, nextEvents[2].EventType)
	assert.True(t, nextEvents[2].Pressed)
	assert.True(t, nextEvents[2].Repeated)
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
