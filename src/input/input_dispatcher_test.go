package input

import (
	"configs"
	"events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetInputDispatcher() (*InputDispatcher, *Keyboard, *TestEmitter) {
	config, _ := configs.NewConfig("testdata/blank.json")
	emitter := NewTestEmitter()
	keyboard := NewKeyboard(emitter)
	return NewInputDispatcher(config, keyboard), keyboard, emitter
}

func Test_NewInputDispatcher(t *testing.T) {
	mapper, _, _ := GetInputDispatcher()
	assert.NotNil(t, mapper, "Mapper failed to initialize")
}

func Test_InputDispatcher_BuildsKeybindingsFromConfig(t *testing.T) {
	config, _ := configs.NewConfig("testdata/keybindings.json")
	emitter := NewTestEmitter()
	keyboard := NewKeyboard(emitter)
	mouse := NewMouse(emitter)
	dispatcher := NewInputDispatcher(config, keyboard, mouse)

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

func Test_IgnoresUnmappedKeyEvents(t *testing.T) {
	_, _, emitter := GetInputDispatcher()
	emitter.fireKeyCallback(KeyQ, KeyPressed)
}

func Test_InputDispatcher_MapsEventsToMultipleKeys(t *testing.T) {
	dispatcher, keyboard, emitter := GetInputDispatcher()
	keyboard.Map("Q", events.Quit)
	keyboard.Map("Esc", events.Quit)

	quitCallCount := 0

	dispatcher.On(events.Quit, func(events.Event) {
		quitCallCount += 1
	})

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyEsc, KeyPressed)

	assert.Equal(t, 2, quitCallCount)
}

func Test_InputDispatcher_MapsKeysToMultipleEvents(t *testing.T) {
	dispatcher, keyboard, emitter := GetInputDispatcher()
	keyboard.Map("J", events.MoveForward)
	keyboard.Map("J", events.MoveBackward)

	forwardCallCount := 0
	backwardCallCount := 0

	dispatcher.On(events.MoveForward, func(events.Event) {
		forwardCallCount += 1
	})

	dispatcher.On(events.MoveBackward, func(events.Event) {
		backwardCallCount += 1
	})

	emitter.fireKeyCallback(KeyJ, KeyPressed)

	assert.Equal(t, 1, forwardCallCount)
	assert.Equal(t, 1, backwardCallCount)
}

func Test_InputDispatcher_ReturnsAndClearsListOfRecentEvents(t *testing.T) {
	dispatcher, keyboard, emitter := GetInputDispatcher()

	keyboard.Map("Q", events.Quit)
	keyboard.Map("D", events.MoveForward)

	emitter.fireKeyCallback(KeyQ, KeyPressed)
	emitter.fireKeyCallback(KeyD, KeyReleased)

	recentEvents := dispatcher.RecentEvents()

	assert.Equal(t, 2, len(recentEvents))
	assert.Equal(t, events.Quit, recentEvents[0].EventType)
	assert.True(t, recentEvents[0].Pressed)

	assert.Equal(t, events.MoveForward, recentEvents[1].EventType)
	assert.False(t, recentEvents[1].Pressed)

	assert.Equal(t, 0, len(dispatcher.RecentEvents()))
}

func Test_InputDispatcher_AsksInputsToPollForCertainEvents(t *testing.T) {
	dispatcher, keyboard, emitter := GetInputDispatcher()
	keyboard.Map("Q", events.Quit)
	keyboard.Map("D", events.MoveForward)
	keyboard.Map("E", events.MoveBackward)

	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
		events.MoveBackward,
	}
	dispatcher.PollEvents(eventList)

	emitter.setKeyState(KeyQ, KeyPressed)
	emitter.setKeyState(KeyD, KeyPressed)
	emitter.setKeyState(KeyE, KeyRepeated)

	nextEvents := dispatcher.RecentEvents()

	assert.Equal(t, 3, len(nextEvents))
	assert.Equal(t, events.Quit, nextEvents[0].EventType)
	assert.Equal(t, events.MoveForward, nextEvents[1].EventType)

	// Repeat key handling
	assert.Equal(t, events.MoveBackward, nextEvents[2].EventType)
	assert.True(t, nextEvents[2].Pressed)
	assert.True(t, nextEvents[2].Repeated)
}

func Test_InputDispatcher_CanStopPollingForSomeEvents(t *testing.T) {
	dispatcher, keyboard, emitter := GetInputDispatcher()
	keyboard.Map("Q", events.Quit)
	keyboard.Map("D", events.MoveForward)

	eventList := []events.EventType{
		events.Quit,
		events.MoveForward,
	}
	dispatcher.PollEvents(eventList)
	dispatcher.UnpollEvents([]events.EventType{events.Quit})

	emitter.setKeyState(KeyQ, KeyPressed)
	emitter.setKeyState(KeyD, KeyPressed)

	nextEvents := dispatcher.RecentEvents()

	assert.Equal(t, 1, len(nextEvents))
	assert.Equal(t, events.MoveForward, nextEvents[0].EventType)
}
