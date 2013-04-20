package input

import (
	"events"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_CanInitalizeAnInput(t *testing.T) {
	mapper := NewInput()
	assert.NotNil(t, mapper, "Mapper failed to initialize")
}

func Test_RegistersACallbackForAnEvent(t *testing.T) {
	mapper := NewInput()
	quitCalled := false

	mapper.On(events.QUIT, func(events.Event) {
		quitCalled = true
	})

	mapper.keyCallback('Q', 1)

	assert.True(t, quitCalled, "Quit callback was not called")
}

func Test_DoesNothingIfNoEventForKey(t *testing.T) {
	mapper := NewInput()

	assert.NotPanics(t, func() {
		mapper.keyCallback('Q', 1)
	}, "Mapper panic'd on unmapped key")
}

func Test_DoesNothingIfNoKeyMappedToEvent(t *testing.T) {
	mapper := NewInput()
	jumpEvent := events.EventType(100)
	jumpCalled := false

	mapper.On(jumpEvent, func(events.Event) {
		jumpCalled = true
	})

	mapper.keyCallback('N', 1)

	assert.False(t, jumpCalled, "Jump event called when it should not have")
}
