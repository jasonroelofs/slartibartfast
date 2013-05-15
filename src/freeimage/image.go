package freeimage

// #cgo LDFLAGS: -lfreeimage
// #include <FreeImage.h>
// #include <stdlib.h>
import "C"
import (
	"log"
	"unsafe"
)

type Image struct {
	ImageType C.FREE_IMAGE_FORMAT

	bitmap *C.FIBITMAP
}

func NewImage(filePath string) *Image {
	image := new(Image)
	image.load(filePath)
	return image
}

func (self *Image) load(filePath string) {
	pathStr := C.CString(filePath)
	defer C.free(unsafe.Pointer(pathStr))

	self.ImageType = C.FreeImage_GetFileType(pathStr, 0)
	self.bitmap = C.FreeImage_Load(self.ImageType, pathStr, 0)

	if self.bitmap == nil {
		log.Panicf("Unable to find image at %s", filePath)
	}
}

func (self *Image) Width() int {
	return int(C.FreeImage_GetWidth(self.bitmap))
}

func (self *Image) Height() int {
	return int(C.FreeImage_GetHeight(self.bitmap))
}

func (self *Image) Unload() {
	C.FreeImage_Unload(self.bitmap)
	self.bitmap = nil
}

func (self *Image) Bytes() *C.BYTE {
	return C.FreeImage_GetBits(self.bitmap)
}
