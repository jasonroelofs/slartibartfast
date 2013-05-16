package render

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func GetMaterialLoader() *MaterialLoader {
	Defaults.LoadPath = "testdata"
	return NewMaterialLoader()
}

func BuildMaterialDef() MaterialDef {
	return MaterialDef{
		Name:    "testMaterial",
		Shaders: "testing",
	}
}

func Test_NetMaterialLoader(t *testing.T) {
	loader := GetMaterialLoader()

	assert.NotNil(t, loader.shaderLoader)
	assert.Equal(t, 0, len(loader.loadedMaterials))
}

func Test_MaterialLoader_Load_TakesMaterialDefReturnsMaterial(t *testing.T) {
	loader := GetMaterialLoader()

	material := loader.Load(BuildMaterialDef())

	assert.IsType(t, Material{}, *material)
	assert.Equal(t, "testMaterial", material.Name)
}

func Test_MaterialLoader_Load_CachesMaterialByNameLocally(t *testing.T) {
	loader := GetMaterialLoader()

	material := loader.Load(BuildMaterialDef())

	assert.Equal(t, material, loader.loadedMaterials["testMaterial"])
}

func Test_MaterialLoader_Load_ComplainsIfMaterialAlreadyLoadedWithSameName(t *testing.T) {
}

func Test_MaterialLoader_Load_LoadsRequestedShaders(t *testing.T) {
	loader := GetMaterialLoader()

	material := loader.Load(BuildMaterialDef())

	assert.Equal(t, "vertex shader\n", material.Shader.Vertex)
	assert.Equal(t, "fragment shader\n", material.Shader.Fragment)

	// To do this here or not. Would require passing in the renderer into all these loaders
	// Might keep this logic inside of Graphical
	//assert.NotNil(t, material.Shader.Program)
}

func Test_MaterialLoader_Get_ReturnsMaterialPointer(t *testing.T) {
	loader := GetMaterialLoader()

	material := loader.Load(BuildMaterialDef())

	assert.Equal(t, material, loader.Get("testMaterial"))
}

func Test_MaterialLoader_Get_ReturnsDefaultMaterialIfNoneMatchName(t *testing.T) {
	loader := GetMaterialLoader()

	defaultMaterial := loader.Load(MaterialDef{
		Name:    "",
		Shaders: "testing",
	})

	assert.Equal(t, defaultMaterial, loader.Get("Some Unknown Material"))
}
