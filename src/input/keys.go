package input

// Custom Key constants, mapped to GLFW
// Here so that we don't have GLFW implementation details spread throughout
// the code base.

import (
	"github.com/go-gl/glfw"
)

// Keyboard characters
const (
	KeyA = 'A' + iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
)

// Special Keys
const (
	KeyEsc = glfw.KeyEsc + iota
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyLshift
	KeyRshift
	KeyLctrl
	KeyRctrl
	KeyLalt
	KeyRalt
	KeyTab
	KeyEnter
	KeyBackspace
	KeyInsert
	KeyDel
	KeyPageup
	KeyPagedown
	KeyHome
	KeyEnd
	KeyKP0
	KeyKP1
	KeyKP2
	KeyKP3
	KeyKP4
	KeyKP5
	KeyKP6
	KeyKP7
	KeyKP8
	KeyKP9
	KeyKPDivide
	KeyKPMultiply
	KeyKPSubtract
	KeyKPAdd
	KeyKPDecimal
	KeyKPEqual
	KeyKPEnter
	KeyKPNumlock
	KeyCapslock
	KeyScrolllock
	KeyPause
	KeyLsuper
	KeyRsuper
	KeyMenu
	KeyLast = KeyMenu
)
