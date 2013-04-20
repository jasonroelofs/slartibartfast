package input

import(
	"events"
	"testing"
)

func Test_CanInitalizeAnInput(t *testing.T) {
	mapper := NewInput()
	if mapper == nil {
		t.Errorf("Mapper failed to initialize")
	}
}

func Test_RegistersACallbackForAnEvent(t *testing.T) {
	mapper := NewInput()
	quitCalled := false

	mapper.On(events.QUIT, func(events.Event) {
		quitCalled = true
	})

	mapper.keyCallback('Q', 1)

	if quitCalled == false {
		t.Errorf("Quit callback was not called")
	}
}

func Test_DoesNothingIfNoEventForKey(t *testing.T) {
	mapper := NewInput()
	mapper.keyCallback('Q', 1)
	// Should not panic
}

func Test_DoesNothingIfNoKeyMappedToEvent(t *testing.T) {
	mapper := NewInput()
	jumpEvent := events.EventType(100)
	jumpCalled := false

	mapper.On(jumpEvent, func(events.Event) {
		jumpCalled = true
	})

	mapper.keyCallback('N', 1)

	if jumpCalled {
		t.Errorf("Jump event should not have been called")
	}
}
