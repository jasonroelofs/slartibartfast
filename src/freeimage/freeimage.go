package freeimage

// #cgo LDFLAGS: -lfreeimage
// #include <FreeImage.h>
// #include <stdlib.h>
import "C"

func Initialise() {
	C.FreeImage_Initialise(0)
}

func DeInitialise() {
	C.FreeImage_DeInitialise()
}

func Version() string {
	return C.GoString(C.FreeImage_GetVersion())
}

func CopyrightMessage() string {
	return C.GoString(C.FreeImage_GetCopyrightMessage())
}
