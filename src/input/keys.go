package input

// Custom Key constants, mapped to GLFW
// Here so that we don't have GLFW implementation details spread throughout
// the code base.

import (
	glfw "github.com/go-gl/glfw3"
)

type KeyCode int

// Keyboard characters
var (
	KeyNone = KeyCode(-1)

	KeyA = defineKey('A', "A")
	KeyB = defineKey('B', "B")
	KeyC = defineKey('C', "C")
	KeyD = defineKey('D', "D")
	KeyE = defineKey('E', "E")
	KeyF = defineKey('F', "F")
	KeyG = defineKey('G', "G")
	KeyH = defineKey('H', "H")
	KeyI = defineKey('I', "I")
	KeyJ = defineKey('J', "J")
	KeyK = defineKey('K', "K")
	KeyL = defineKey('L', "L")
	KeyM = defineKey('M', "M")
	KeyN = defineKey('N', "N")
	KeyO = defineKey('O', "O")
	KeyP = defineKey('P', "P")
	KeyQ = defineKey('Q', "Q")
	KeyR = defineKey('R', "R")
	KeyS = defineKey('S', "S")
	KeyT = defineKey('T', "T")
	KeyU = defineKey('U', "U")
	KeyV = defineKey('V', "V")
	KeyW = defineKey('W', "W")
	KeyX = defineKey('X', "X")
	KeyY = defineKey('Y', "Y")
	KeyZ = defineKey('Z', "Z")
	Key0 = defineKey('0', "0")
	Key1 = defineKey('1', "1")
	Key2 = defineKey('2', "2")
	Key3 = defineKey('3', "3")
	Key4 = defineKey('4', "4")
	Key5 = defineKey('5', "5")
	Key6 = defineKey('6', "6")
	Key7 = defineKey('7', "7")
	Key8 = defineKey('8', "8")
	Key9 = defineKey('9', "9")
)

// Special Keys
var (
	KeyEsc          = defineKey(glfw.KeyEscape, "Esc")
	KeyF1           = defineKey(glfw.KeyF1, "F1")
	KeyF2           = defineKey(glfw.KeyF2, "F2")
	KeyF3           = defineKey(glfw.KeyF3, "F3")
	KeyF4           = defineKey(glfw.KeyF4, "F4")
	KeyF5           = defineKey(glfw.KeyF5, "F5")
	KeyF6           = defineKey(glfw.KeyF6, "F6")
	KeyF7           = defineKey(glfw.KeyF7, "F7")
	KeyF8           = defineKey(glfw.KeyF8, "F8")
	KeyF9           = defineKey(glfw.KeyF9, "F9")
	KeyF10          = defineKey(glfw.KeyF10, "F10")
	KeyF11          = defineKey(glfw.KeyF11, "F11")
	KeyF12          = defineKey(glfw.KeyF12, "F12")
	KeyF13          = defineKey(glfw.KeyF13, "F13")
	KeyF14          = defineKey(glfw.KeyF14, "F14")
	KeyF15          = defineKey(glfw.KeyF15, "F15")
	KeyF16          = defineKey(glfw.KeyF16, "F16")
	KeyF17          = defineKey(glfw.KeyF17, "F17")
	KeyF18          = defineKey(glfw.KeyF18, "F18")
	KeyF19          = defineKey(glfw.KeyF19, "F19")
	KeyF20          = defineKey(glfw.KeyF20, "F20")
	KeyF21          = defineKey(glfw.KeyF21, "F21")
	KeyF22          = defineKey(glfw.KeyF22, "F22")
	KeyF23          = defineKey(glfw.KeyF23, "F23")
	KeyF24          = defineKey(glfw.KeyF24, "F24")
	KeyF25          = defineKey(glfw.KeyF25, "F25")
	KeyUp           = defineKey(glfw.KeyUp, "Up")
	KeyDown         = defineKey(glfw.KeyDown, "Down")
	KeyLeft         = defineKey(glfw.KeyLeft, "Left")
	KeyRight        = defineKey(glfw.KeyRight, "Right")
	KeyLeftShift    = defineKey(glfw.KeyLeftShift, "LeftShift")
	KeyRightShift   = defineKey(glfw.KeyRightShift, "RightShift")
	KeyLeftControl  = defineKey(glfw.KeyLeftControl, "LeftControl")
	KeyRightControl = defineKey(glfw.KeyRightControl, "RightControl")
	KeyLeftAlt      = defineKey(glfw.KeyLeftAlt, "LeftAlt")
	KeyRightAlt     = defineKey(glfw.KeyRightAlt, "RightAlt")
	KeyTab          = defineKey(glfw.KeyTab, "Tab")
	KeyEnter        = defineKey(glfw.KeyEnter, "Enter")
	KeyBackspace    = defineKey(glfw.KeyBackspace, "Backspace")
	KeyInsert       = defineKey(glfw.KeyInsert, "Insert")
	KeyDelete       = defineKey(glfw.KeyDelete, "Delete")
	KeyPageUp       = defineKey(glfw.KeyPageUp, "PageUp")
	KeyPageDown     = defineKey(glfw.KeyPageDown, "PageDown")
	KeyHome         = defineKey(glfw.KeyHome, "Home")
	KeyEnd          = defineKey(glfw.KeyEnd, "End")
	KeyKP0          = defineKey(glfw.KeyKp0, "KP0")
	KeyKP1          = defineKey(glfw.KeyKp1, "KP1")
	KeyKP2          = defineKey(glfw.KeyKp2, "KP2")
	KeyKP3          = defineKey(glfw.KeyKp3, "KP3")
	KeyKP4          = defineKey(glfw.KeyKp4, "KP4")
	KeyKP5          = defineKey(glfw.KeyKp5, "KP5")
	KeyKP6          = defineKey(glfw.KeyKp6, "KP6")
	KeyKP7          = defineKey(glfw.KeyKp7, "KP7")
	KeyKP8          = defineKey(glfw.KeyKp8, "KP8")
	KeyKP9          = defineKey(glfw.KeyKp9, "KP9")
	KeyKPDivide     = defineKey(glfw.KeyKpDivide, "KPDivide")
	KeyKPMultiply   = defineKey(glfw.KeyKpMultiply, "KPMultiply")
	KeyKPSubtract   = defineKey(glfw.KeyKpSubtract, "KPSubtract")
	KeyKPAdd        = defineKey(glfw.KeyKpAdd, "KPAdd")
	KeyKPDecimal    = defineKey(glfw.KeyKpDecimal, "KPDecimal")
	KeyKPEqual      = defineKey(glfw.KeyKpEqual, "KPEqual")
	KeyKPEnter      = defineKey(glfw.KeyKpEnter, "KPEnter")
	KeyCapsLock     = defineKey(glfw.KeyCapsLock, "CapsLock")
	KeyScrollLock   = defineKey(glfw.KeyScrollLock, "ScrollLock")
	KeyPause        = defineKey(glfw.KeyPause, "Pause")
	KeyLeftSuper    = defineKey(glfw.KeyLeftSuper, "LeftSuper")
	KeyRightSuper   = defineKey(glfw.KeyRightSuper, "RightSuper")
	KeyMenu         = defineKey(glfw.KeyMenu, "Menu")
	KeySpace        = defineKey(glfw.KeySpace, "Space")
	KeyApostrophe   = defineKey(glfw.KeyApostrophe, "'")
	KeyComma        = defineKey(glfw.KeyComma, ",")
	KeyMinus        = defineKey(glfw.KeyMinus, "-")
	KeyPeriod       = defineKey(glfw.KeyPeriod, ".")
	KeySlash        = defineKey(glfw.KeySlash, "/")
	KeyBackslash    = defineKey(glfw.KeyBackslash, "\\")
	KeySemicolon    = defineKey(glfw.KeySemicolon, ";")
	KeyEqual        = defineKey(glfw.KeyEqual, "=")
	KeyLeftBracket  = defineKey(glfw.KeyLeftBracket, "[")
	KeyRightBracket = defineKey(glfw.KeyRightBracket, "]")
	KeyGraveAccent  = defineKey(glfw.KeyGraveAccent, "`")
	KeyLast         = glfw.KeyLast
)

var ()

// Returns a KeyCode for the given key name.
func KeyFromName(keyName string) KeyCode {
	code, ok := keyNameToCode[keyName]
	if ok {
		return code
	} else {
		return KeyNone
	}
}

var (
	keyNameToCode = make(map[string]KeyCode)
)

func defineKey(keyCode glfw.Key, keyName string) KeyCode {
	code := KeyCode(keyCode)
	keyNameToCode[keyName] = code
	return code
}
