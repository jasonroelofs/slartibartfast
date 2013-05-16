package freeimage

// #cgo LDFLAGS: -lfreeimage
// #include <FreeImage.h>
// #include <stdlib.h>
import "C"

var initialised = false

func Initialise() {
	C.FreeImage_Initialise(0)
	initialised = true
}

func DeInitialise() {
	C.FreeImage_DeInitialise()
	initialised = false
}

func Version() string {
	return C.GoString(C.FreeImage_GetVersion())
}

func CopyrightMessage() string {
	return C.GoString(C.FreeImage_GetCopyrightMessage())
}
