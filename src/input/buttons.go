package input

// Custom Button constants, mapped to GLFW
// Here so that we don't have GLFW implementation details spread throughout
// the code base.

import (
	glfw "github.com/go-gl/glfw3"
)

// Mouse buttons
const (
	Mouse1 = glfw.MouseButton1 + iota
	Mouse2
	Mouse3
	Mouse4
	Mouse5
	Mouse6
	Mouse7
	Mouse8
	MouseLast   = glfw.MouseButtonLast
	MouseLeft   = glfw.MouseButtonLeft
	MouseRight  = glfw.MouseButtonRight
	MouseMiddle = glfw.MouseButtonMiddle
)
