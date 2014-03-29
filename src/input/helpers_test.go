package input

import (
	"events"
)

// Bogus InputEmitter object for use in tests

type TestEmitter struct {
	// Implements input.InputEmitter

	keyCallback      func(KeyCode, KeyState)
	mousePosCallback func(int, int)

	mouseButtonCallback func(MouseButtonCode, KeyState)

	keyStates map[KeyCode]KeyState
}

func NewTestEmitter() *TestEmitter {
	return &TestEmitter{
		keyStates: make(map[KeyCode]KeyState),
	}
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

func (self *TestEmitter) fireMousePositionCallback(x, y int) {
	self.mousePosCallback(x, y)
}

func (self *TestEmitter) ShowCursor() {
}

func (self *TestEmitter) HideCursor() {
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

// Bogus InputQueue object

type TestInputListener struct {
	ReceivedEvents EventList
}

func NewListener() *TestInputListener {
	return &TestInputListener{}
}

func (self *TestInputListener) ReceiveEvent(event events.Event) {
	self.ReceivedEvents = append(self.ReceivedEvents, event)
}
