package input

// Custom Button constants, mapped to GLFW
// Here so that we don't have GLFW implementation details spread throughout
// the code base.

import (
	"github.com/go-gl/glfw"
)

// Mouse buttons
const (
	Mouse1 = glfw.Mouse1 + iota
	Mouse2
	Mouse3
	Mouse4
	Mouse5
	Mouse6
	Mouse7
	Mouse8
	MouseLast   = Mouse8
	MouseLeft   = Mouse1
	MouseRight  = Mouse2
	MouseMiddle = Mouse3
)
