package render

import (
	"github.com/stretchr/testify/assert"
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
