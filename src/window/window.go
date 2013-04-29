package window

import (
	"configs"
	"errors"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
)

type WindowConfig struct {
	Width      uint
	Height     uint
	Fullscreen bool
}

func Open(config *configs.Config) {
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	windowConfig := WindowConfig{}
	err = config.Get("window", &windowConfig)
	if err != nil {
		panic(err)
	}

	windowFlags := glfw.Windowed
	if windowConfig.Fullscreen {
		windowFlags = glfw.Fullscreen
	}

	// Force OpenGL 3.2
	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 2)
	glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.OpenWindowHint(glfw.OpenGLForwardCompat, gl.TRUE)

	err = glfw.OpenWindow(
		int(windowConfig.Width), int(windowConfig.Height),
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

func StillOpen() bool {
	return glfw.WindowParam(glfw.Opened) == gl.TRUE
}

func Present() {
	glfw.SwapBuffers()
}

func Close() {
	glfw.Terminate()
}
