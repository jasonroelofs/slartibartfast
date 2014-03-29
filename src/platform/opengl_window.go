package platform

import (
	"configs"
	"errors"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"input"
	"log"
)

type OpenGLWindow struct {
	// Implements core.Window
	// Implements input.InputEmitter

	config windowConfig
	window *glfw.Window

	timeLastCall float64
	hiddenCursor bool
}

type windowConfig struct {
	Width      uint
	Height     uint
	Fullscreen bool
	VSync      bool
}

func NewOpenGLWindow(config *configs.Config) *OpenGLWindow {
	glWindow := new(OpenGLWindow)

	windowConfig := windowConfig{}
	err := config.Get("window", &windowConfig)
	if err != nil {
		panic(err)
	}

	glWindow.config = windowConfig

	return glWindow
}

func (self *OpenGLWindow) Open() {
	var err error
	if !glfw.Init() {
		panic(errors.New("Unable to initialize GLFW"))
	}

	glfw.SetErrorCallback(func(code glfw.ErrorCode, desc string) {
		log.Printf("[GLFW Error] (%d) %s", code, desc)
	})

	var monitor *glfw.Monitor
	if self.config.Fullscreen {
		monitor, err = glfw.GetPrimaryMonitor()

		if err != nil {
			panic(err)
		}
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)

	// Default buffer sizes
	glfw.WindowHint(glfw.DepthBits, 32)
	glfw.WindowHint(glfw.StencilBits, 0)

	// Toggle VSync. Turning VSync off aparently doesn't work via glfw through
	// some ATI cards and drivers
	if self.config.VSync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	self.window, err = glfw.CreateWindow(
		int(self.config.Width), int(self.config.Height),
		"Project Slartibartfast",
		monitor,
		nil)

	if err != nil {
		panic(err)
	}

	self.window.MakeContextCurrent()

	if glewError := gl.Init(); glewError != 0 {
		panic(errors.New("Unable to initialize OpenGL"))
	}
}

func (self *OpenGLWindow) AspectRatio() float32 {
	width, height := self.window.GetSize()
	return float32(width) / float32(height)
}

func (self *OpenGLWindow) TimeSinceLast() float32 {
	if self.timeLastCall == 0 {
		self.timeLastCall = glfw.GetTime()
	}

	now := glfw.GetTime()
	diff := now - self.timeLastCall
	self.timeLastCall = now

	return float32(diff)
}

func (self *OpenGLWindow) IsOpen() bool {
	return !self.window.ShouldClose()
}

func (self *OpenGLWindow) SwapBuffers() {
	self.window.SwapBuffers()
	glfw.PollEvents()
}

func (self *OpenGLWindow) Close() {
	self.window.Destroy()
	glfw.Terminate()
}

// KeyCallback :: input.InputEmitter
func (self *OpenGLWindow) KeyCallback(callback func(input.KeyCode, input.KeyState)) {
	self.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, state glfw.Action, mods glfw.ModifierKey) {
		callback(input.KeyCode(key), input.KeyState(state))
	})
}

// MouseButtonCallback :: input.InputEmitter
func (self *OpenGLWindow) MouseButtonCallback(callback func(input.MouseButtonCode, input.KeyState)) {
	self.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, state glfw.Action, mod glfw.ModifierKey) {
		callback(input.MouseButtonCode(button), input.KeyState(state))
	})
}

// MousePositionCallback :: input.InputEmitter
func (self *OpenGLWindow) MousePositionCallback(callback func(int, int)) {
	self.window.SetCursorPositionCallback(func(w *glfw.Window, xIn, yIn float64) {
		x := int(xIn)
		y := int(yIn)

		if self.hiddenCursor {
			// Don't care about actual diff from center, we just want the distance travelled
			// for this singular event. We always reset back to origin. Use for FPS-type controls
			callback(x, y)
			self.resetCursor()
		} else {
			// GLFW puts 0,0 at the top left of the window. Need to transform this
			// origin to be the center of the screen.
			windowX, windowY := self.window.GetSize()
			xPos := x - (windowX / 2)
			yPos := (windowY / 2) - y

			callback(xPos, yPos)
		}
	})
}

// ShowCursor :: input.InputEmitter
func (self *OpenGLWindow) ShowCursor() {
	self.hiddenCursor = false
	self.window.SetInputMode(glfw.Cursor, glfw.CursorNormal)
}

// HideCursor :: input.InputEmitter
func (self *OpenGLWindow) HideCursor() {
	self.hiddenCursor = true
	self.window.SetInputMode(glfw.Cursor, glfw.CursorDisabled)
	self.resetCursor()
}

func (self *OpenGLWindow) resetCursor() {
	self.window.SetCursorPosition(0, 0)
}

// MouseWheelCallback :: input.InputEmitter
func (self *OpenGLWindow) MouseWheelCallback(callback func(int)) {
	self.window.SetScrollCallback(func(w *glfw.Window, xoff, yoff float64) {
		callback(int(xoff))
	})
}

// IsKeyPressed :: input.InputEmitter
func (self *OpenGLWindow) IsKeyPressed(key input.KeyCode) bool {
	keyState := self.window.GetKey(glfw.Key(key))
	return keyState == glfw.Press || keyState == glfw.Repeat
}

// IsKeyRepeated :: input.InputEmitter
func (self *OpenGLWindow) IsKeyRepeated(key input.KeyCode) bool {
	keyState := self.window.GetKey(glfw.Key(key))
	return keyState == glfw.Repeat
}
