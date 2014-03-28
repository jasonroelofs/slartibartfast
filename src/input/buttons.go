package input

// Custom Button constants, mapped to GLFW
// Here so that we don't have GLFW implementation details spread throughout
// the code base.

import (
	glfw "github.com/go-gl/glfw3"
)

type MouseButtonCode int

// Mouse buttons
var (
	MouseNone = MouseButtonCode(-1)

	Mouse1 			= defineMouseButton(glfw.MouseButton1, "Mouse1")
	Mouse2 			= defineMouseButton(glfw.MouseButton2, "Mouse2")
	Mouse3 			= defineMouseButton(glfw.MouseButton3, "Mouse3")
	Mouse4 			= defineMouseButton(glfw.MouseButton4, "Mouse4")
	Mouse5 			= defineMouseButton(glfw.MouseButton5, "Mouse5")
	Mouse6 			= defineMouseButton(glfw.MouseButton6, "Mouse6")
	Mouse7 			= defineMouseButton(glfw.MouseButton7, "Mouse7")
	Mouse8 			= defineMouseButton(glfw.MouseButton8, "Mouse8")
	MouseLast   = defineMouseButton(glfw.MouseButtonLast, "MouseLast")
	MouseLeft   = defineMouseButton(glfw.MouseButtonLeft, "MouseLeft")
	MouseRight  = defineMouseButton(glfw.MouseButtonRight, "MouseRight")
	MouseMiddle = defineMouseButton(glfw.MouseButtonMiddle, "MouseMiddle")
)

// Returns a MouseButtonCode for the given mouseButton name.
func MouseButtonFromName(mouseButtonName string) MouseButtonCode {
	code, ok := mouseButtonNameToCode[mouseButtonName]
	if ok {
		return code
	} else {
		return MouseNone
	}
}

var (
	mouseButtonNameToCode = make(map[string]MouseButtonCode)
)

func defineMouseButton(buttonCode glfw.MouseButton, buttonName string) MouseButtonCode {
	code := MouseButtonCode(buttonCode)
	mouseButtonNameToCode[buttonName] = code
	return code
}
