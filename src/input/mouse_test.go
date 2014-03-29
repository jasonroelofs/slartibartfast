package input

import (
	"events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Mouse_MapsMovementToEvent(t *testing.T) {
	emitter := NewTestEmitter()
	listener := NewListener()
	mouse := NewMouse(emitter)

	mouse.OnEvent(listener.ReceiveEvent)

	emitter.fireMousePositionCallback(1, 2)
	emitter.fireMousePositionCallback(3, 4)

	assert.Equal(t, 2, len(listener.ReceivedEvents))
	assert.Equal(t, events.MouseMove, listener.ReceivedEvents[0].EventType)
	assert.Equal(t, 1, listener.ReceivedEvents[0].MouseXDiff)
	assert.Equal(t, 2, listener.ReceivedEvents[0].MouseYDiff)

	assert.Equal(t, events.MouseMove, listener.ReceivedEvents[1].EventType)
	assert.Equal(t, 3, listener.ReceivedEvents[1].MouseXDiff)
	assert.Equal(t, 4, listener.ReceivedEvents[1].MouseYDiff)
}

func Test_Mouse_MapsButtonsToEvents(t *testing.T) {
	emitter := NewTestEmitter()
	listener := NewListener()
	mouse := NewMouse(emitter)

	mouse.OnEvent(listener.ReceiveEvent)

	mouse.Map("Mouse1", events.MoveForward)
	mouse.Map("Mouse2", events.MoveBackward)

	emitter.fireMouseButtonCallback(Mouse1, KeyPressed)
	emitter.fireMouseButtonCallback(Mouse2, KeyPressed)

	assert.Equal(t, 2, len(listener.ReceivedEvents))
	assert.Equal(t, events.MoveForward, listener.ReceivedEvents[0].EventType)
	assert.True(t, listener.ReceivedEvents[0].Pressed)

	assert.Equal(t, events.MoveBackward, listener.ReceivedEvents[1].EventType)
	assert.True(t, listener.ReceivedEvents[1].Pressed)
}

