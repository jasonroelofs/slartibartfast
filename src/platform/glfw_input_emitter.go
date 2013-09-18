package platform

import (
	"github.com/go-gl/glfw"
	"input"
)

// GLFWInputEmitter implements input emission using GLFW's input handling
type GLFWInputEmitter struct {
	// Implements input.InputEmitter
}

func (self *GLFWInputEmitter) KeyCallback(callback func(input.KeyCode, input.KeyState)) {
	glfw.SetKeyCallback(func(key, state int) {
		callback(input.KeyCode(key), input.KeyState(state))
	})
}

func (self *GLFWInputEmitter) MouseButtonCallback(callback func(input.KeyCode, input.KeyState)) {
	glfw.SetMouseButtonCallback(func(key, state int) {
		callback(input.KeyCode(key), input.KeyState(state))
	})
}

func (self *GLFWInputEmitter) MousePositionCallback(callback func(int, int)) {
	// GLFW puts 0,0 at the top left of the window. Need to transform this
	// origin to be the center of the screen.
	windowX, windowY := glfw.WindowSize()
	glfw.SetMousePosCallback(func(x, y int) {
		callback(x - (windowX / 2), (windowY / 2) - y)
	})
}

func (self *GLFWInputEmitter) MouseWheelCallback(callback func(int)) {
	// TODO Calculate the distance from the last scroll wheel setting.
	// This is changed in GLFW 3 to always be last-moved distance
	glfw.SetMouseWheelCallback(callback)
}

func (self *GLFWInputEmitter) IsKeyPressed(key input.KeyCode) bool {
	return glfw.Key(int(key)) == glfw.KeyPress
}
