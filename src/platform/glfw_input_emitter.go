package platform

import (
	"github.com/go-gl/glfw"
	"input"
)

// GLFWInputEmitter implements input emission using GLFW's input handling
type GLFWInputEmitter struct {
	// Implements input.InputEmitter

	hiddenCursor bool
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
	glfw.SetMousePosCallback(func(x, y int) {
		if self.hiddenCursor {
			// Don't care about actual diff from center, we just want the distance travelled
			// for this singular event. We always reset back to origin. Use for FPS-type controls
			callback(x, y)
			self.resetCursor()
		} else {
			// GLFW puts 0,0 at the top left of the window. Need to transform this
			// origin to be the center of the screen.
			windowX, windowY := glfw.WindowSize()
			xPos := x - (windowX / 2)
			yPos := (windowY / 2) - y

			callback(xPos, yPos)
		}
	})
}

func (self *GLFWInputEmitter) ShowCursor() {
	self.hiddenCursor = false
	glfw.Enable(glfw.MouseCursor)
}

func (self *GLFWInputEmitter) HideCursor() {
	self.hiddenCursor = true
	glfw.Disable(glfw.MouseCursor)
	self.resetCursor()
}

func (self *GLFWInputEmitter) resetCursor() {
	glfw.SetMousePos(0, 0)
}

func (self *GLFWInputEmitter) MouseWheelCallback(callback func(int)) {
	// TODO Calculate the distance from the last scroll wheel setting.
	// This is changed in GLFW 3 to always be last-moved distance
	glfw.SetMouseWheelCallback(callback)
}

func (self *GLFWInputEmitter) IsKeyPressed(key input.KeyCode) bool {
	return glfw.Key(int(key)) == glfw.KeyPress
}
