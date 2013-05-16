package render

import (
	"freeimage"
)

type Texture struct {
	Image *freeimage.Image

	// Link to the loaded entry on the GPU
	Id interface{}
}
