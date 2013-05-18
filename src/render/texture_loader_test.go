package render

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewTextureLoader(t *testing.T) {
	loader := NewTextureLoader()
	assert.Equal(t, 0, len(loader.loadedTextures))
}

func Test_TextureLoader_Load(t *testing.T) {
	Defaults.LoadPath = "testdata"
	loader := NewTextureLoader()

	texture := loader.Load("test.png")

	assert.Equal(t, 200, texture.Image.Width())
	assert.Equal(t, 200, texture.Image.Height())
}

func Test_TextureLoader_Load_CachesLoadedTexturesByName(t *testing.T) {
	Defaults.LoadPath = "testdata"
	loader := NewTextureLoader()

	texture := loader.Load("test.png")

	assert.Equal(t, texture, loader.loadedTextures["test.png"])
}
