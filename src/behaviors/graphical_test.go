package behaviors

import (
	"components"
	"core"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

type TestRenderer struct {
	loadedMesh     *core.Mesh
	loadedMaterial *core.Material
}

func (self *TestRenderer) LoadMesh(mesh *core.Mesh) {
	self.loadedMesh = mesh
}

func (self *TestRenderer) LoadMaterial(material *core.Material) {
	self.loadedMaterial = material
}

func (self *TestRenderer) BeginRender() {
}

func (self *TestRenderer) Render(mesh *core.Mesh, material *core.Material) {
}

func (self *TestRenderer) FinishRender() {
}

func getTestGraphical() (*Graphical, *TestRenderer, *core.EntityDB) {
	entityDB := new(core.EntityDB)
	renderer := new(TestRenderer)

	defaults.LoadPath = "testdata"
	graphical := NewGraphical(renderer, entityDB)

	return graphical, renderer, entityDB
}

func Test_NewGraphical(t *testing.T) {
	graphical, renderer, _ := getTestGraphical()

	assert.NotNil(t, graphical)
	assert.NotNil(t, graphical.entitySet)
	assert.Equal(t, renderer, graphical.renderer)
}

func Test_NewGraphical_PreLoadsDefaultMesh(t *testing.T) {
	_, renderer, _ := getTestGraphical()

	assert.Equal(t, core.DefaultMesh, renderer.loadedMesh)
}

func Test_NewGraphical_PreLoadsDefaultMaterial(t *testing.T) {
	_, renderer, _ := getTestGraphical()

	assert.Equal(t, core.DefaultMaterial, renderer.loadedMaterial)
}

func Test_LoadMaterial_ReadsShadersAndSavesCodeToMaterial(t *testing.T) {
	getTestGraphical()

	// Check shader code of default material, loaded by NewGraphical
	assert.Equal(t, core.DefaultMaterial.VertexShader, "vertex shader\n")
	assert.Equal(t, core.DefaultMaterial.FragmentShader, "fragment shader\n")
}

func Test_LoadMaterial_IgnoresFileLoadIfShaderSourceExists(t *testing.T) {
	core.DefaultMaterial.VertexShader = "default vertex shader"
	core.DefaultMaterial.FragmentShader = "default fragment shader"

	getTestGraphical()

	// Check shader code of default material, loaded by NewGraphical
	assert.Equal(t, core.DefaultMaterial.VertexShader, "default vertex shader")
	assert.Equal(t, core.DefaultMaterial.FragmentShader, "default fragment shader")
}

func Test_LoadMaterial_FallsBackToDefaultMeshIfNoShaderFilesFound(t *testing.T) {
	core.DefaultMaterial.VertexShader = "default vertex shader"
	core.DefaultMaterial.FragmentShader = "default fragment shader"

	material := &core.Material{
		Name: "testMat",
		Shaders: "missing",
	}

	graphical, _, _ := getTestGraphical()
	graphical.LoadMaterial(material)

	assert.Equal(t, material.Name, core.DefaultMaterial.Name)
	assert.Equal(t, core.DefaultMaterial.VertexShader, material.VertexShader)
	assert.Equal(t, core.DefaultMaterial.FragmentShader, material.FragmentShader)
}

func Test_SetUpEntity_SetsDefaultMeshIfNoneSpecified(t *testing.T) {
	graphical, _, _ := getTestGraphical()

	entity := core.NewEntity()
	entity.AddComponent(new(components.Visual))

	graphical.SetUpEntity(entity)

	visual := entity.GetComponent(components.VISUAL).(*components.Visual)
	assert.Equal(t, core.DefaultMesh.Name, visual.MeshName)
}

func Test_SetUpEntity_LinksLoadedMeshIfNameMatches(t *testing.T) {
}

func Test_SetUpEntity_TellsRendererToLoadNewMesh(t *testing.T) {
}

func Test_SetUpEntity_SetsDefaultMaterialIfNoneSpecified(t *testing.T) {
	graphical, _, _ := getTestGraphical()

	entity := core.NewEntity()
	entity.AddComponent(new(components.Visual))

	graphical.SetUpEntity(entity)

	visual := entity.GetComponent(components.VISUAL).(*components.Visual)
	assert.Equal(t, core.DefaultMaterial.Name, visual.MaterialName)
}

func Test_SetUpEntity_LinksLoadedMaterialIfNameMatches(t *testing.T) {
}

func Test_SetUpEntity_TellsRendererToLoadNewMaterial(t *testing.T) {
}

func Test_Update_RendersAllVisualEntityMeshesWithMaterials(t *testing.T) {
}
