package render

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewShaderLoader(t *testing.T) {
	loader := NewShaderLoader()
	assert.Equal(t, 0, len(loader.loadedShaders))
}

func Test_ShaderLoader_Load(t *testing.T) {
	Defaults.LoadPath = "testdata"
	loader := NewShaderLoader()

	shader := loader.Load("testing")

	assert.Equal(t, "vertex shader\n", shader.Vertex)
	assert.Equal(t, "fragment shader\n", shader.Fragment)
}

func Test_ShaderLoader_Load_CachesLoadedShadersByName(t *testing.T) {
	Defaults.LoadPath = "testdata"
	loader := NewShaderLoader()

	shader := loader.Load("testing")

	assert.Equal(t, shader, loader.loadedShaders["testing"])
}

func Test_ShaderLoader_Load_PanicsIfCantFindShaderFile(t *testing.T) {
	Defaults.LoadPath = "testdata"
	loader := NewShaderLoader()

	assert.Panics(t, func() {
		loader.Load("does_not_exist")
	})
}
