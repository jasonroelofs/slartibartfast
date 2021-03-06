package input

// InputEmitter emits input events. This is mainly through callbacks according
// to the device emitting the event, whether it be a keyboard, mouse, or joystick.
type InputEmitter interface {
	// KeyCallback sets a callback that fires when something happens with a key
	// The callback is passed two parameters: key and state
	// key is the Key pressed (see keys.go)
	// state is one of the KeyState constants
	KeyCallback(func(KeyCode, KeyState))

	// MouseButtonCallback works like KeyCallback but for Mouse buttons
	MouseButtonCallback(func(MouseButtonCode, KeyState))

	// MousePositionCallback sets a callback that fires when the mouse is moved
	// The callback is passed the X and Y pixel position of the mouse cursor's new position
	// relative to the center of the screen.
	MousePositionCallback(func(int, int))

	// MouseWheelCallback sets a callback that fires when the mouse wheel is scrolled
	// The callback is passed the distance the scroll wheel moved
	MouseWheelCallback(func(int))

	// ShowCursor flags the system to show the mouse cursor
	ShowCursor()

	// HideCursor flags the system to hide the mouse cursor
	HideCursor()

	// IsKeyPressed returns true or false depending on if the given Key is currently depressed
	IsKeyPressed(KeyCode) bool

	// IsKeyRepeated returns true or false depending on if the given Key is currently marked as repeating
	IsKeyRepeated(KeyCode) bool
}
