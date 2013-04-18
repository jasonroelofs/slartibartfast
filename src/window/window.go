package window

import (
	"github.com/go-gl/glfw"
	"time"
)

func NewWindow() {
	glfw.Init()
	defer glfw.Terminate()

	glfw.OpenWindow(648, 480, 0, 0, 0, 0, 0, 0, glfw.Windowed)

	time.Sleep(10 * time.Second)
}
