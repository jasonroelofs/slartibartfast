package platform

import (
	"configs"
	"errors"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
)

type OpenGLWindow struct {
	// Implements core.Window
	config windowConfig
}

type windowConfig struct {
	Width      uint
	Height     uint
	Fullscreen bool
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
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	windowFlags := glfw.Windowed
	if self.config.Fullscreen {
		windowFlags = glfw.Fullscreen
	}

	// Force OpenGL 3.2
	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 2)
	glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	err = glfw.OpenWindow(
		int(self.config.Width), int(self.config.Height),
		// r, g, b, a
		0, 0, 0, 0,
		// depth, stencil
		32, 0,
		windowFlags)

	if err != nil {
		panic(errors.New("Unable to open window!"))
	}

	if glewError := gl.Init(); glewError != 0 {
		panic(errors.New("Unable to initialize OpenGL"))
	}

	glfw.SetWindowTitle("Project Slartibartfast")
}

func (self *OpenGLWindow) IsOpen() bool {
	return glfw.WindowParam(glfw.Opened) == gl.TRUE
}

func (self *OpenGLWindow) SwapBuffers() {
	glfw.SwapBuffers()
}

func (self *OpenGLWindow) Close() {
	glfw.Terminate()
}
