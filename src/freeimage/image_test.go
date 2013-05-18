package freeimage

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewImage_InitializesImageWithLoadedData(t *testing.T) {
	image := NewImage("testdata/test.png")
	assert.NotNil(t, image)
	assert.NotNil(t, image.bitmap)

	assert.Equal(t, PNG, image.ImageType)
}

func Test_Image_Width(t *testing.T) {
	image := NewImage("testdata/test.png")
	assert.Equal(t, 200, image.Width())
}

func Test_Image_Height(t *testing.T) {
	image := NewImage("testdata/test.png")
	assert.Equal(t, 200, image.Height())
}

func Test_Image_Unload(t *testing.T) {
	image := NewImage("testdata/test.png")
	image.Unload()
	assert.Nil(t, image.bitmap)
}

func Test_Image_Bytes(t *testing.T) {
	image := NewImage("testdata/test.png")
	assert.NotNil(t, image.Bytes())
}
