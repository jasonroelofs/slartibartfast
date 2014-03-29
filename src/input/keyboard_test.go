package input

import (
	"events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Keyboard_MapsKeysToEvents(t *testing.T) {
	emitter := NewTestEmitter()
	listener := NewListener()
	keyboard := NewKeyboard(emitter)

	keyboard.OnEvent(listener.ReceiveEvent)

	keyboard.Map("E", events.MoveForward)
	keyboard.Map("D", events.MoveBackward)

	emitter.fireKeyCallback(KeyE, KeyPressed)
	emitter.fireKeyCallback(KeyD, KeyRepeated)
	emitter.fireKeyCallback(KeyF, KeyPressed)
	emitter.fireKeyCallback(KeyE, KeyReleased)

	assert.Equal(t, 3, len(listener.ReceivedEvents))
	assert.Equal(t, events.MoveForward, listener.ReceivedEvents[0].EventType)
	assert.True(t, listener.ReceivedEvents[0].Pressed)

	assert.Equal(t, events.MoveBackward, listener.ReceivedEvents[1].EventType)
	assert.True(t, listener.ReceivedEvents[1].Pressed)
	assert.True(t, listener.ReceivedEvents[1].Repeated)

	assert.Equal(t, events.MoveForward, listener.ReceivedEvents[2].EventType)
	assert.False(t, listener.ReceivedEvents[2].Pressed)
}

func Test_Keyboard_OnKey_FiresOnSpecificKey(t *testing.T) {
	emitter := NewTestEmitter()
	listener := NewListener()
	keyboard := NewKeyboard(emitter)

	keyboard.OnEvent(listener.ReceiveEvent)

	pKeyHit := false
	var pKeyEvent events.Event

	keyboard.OnKey(KeyP, func(event events.Event) {
		pKeyHit = true
		pKeyEvent = event
	})

	emitter.fireKeyCallback(KeyP, KeyPressed)

	assert.True(t, pKeyHit, "P key callback was not called")
	assert.True(t, pKeyEvent.Pressed)
}
