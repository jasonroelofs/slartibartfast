package behaviors

import (
	"components"
	"core"
	"github.com/stretchrcom/testify/assert"
	"math3d"
	"render"
	"testing"
)

type TestRenderer struct {
	loadedMesh     *render.Mesh
	loadedMaterial *render.Material
	queueRendered  *render.RenderQueue
}

func (self *TestRenderer) LoadMesh(mesh *render.Mesh) {
	self.loadedMesh = mesh
}

func (self *TestRenderer) LoadMaterial(material *render.Material) {
	self.loadedMaterial = material
}

func (self *TestRenderer) BeginRender() {
}

func (self *TestRenderer) Render(queue *render.RenderQueue) {
	self.queueRendered = queue
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

	assert.Equal(t, render.DefaultMesh, renderer.loadedMesh)
}

func Test_NewGraphical_PreLoadsDefaultMaterial(t *testing.T) {
	_, renderer, _ := getTestGraphical()

	assert.Equal(t, render.DefaultMaterial, renderer.loadedMaterial)
}

func Test_LoadMaterial_ReadsShadersAndSavesCodeToMaterial(t *testing.T) {
	getTestGraphical()

	// Check shader code of default material, loaded by NewGraphical
	assert.Equal(t, render.DefaultMaterial.VertexShader, "vertex shader\n")
	assert.Equal(t, render.DefaultMaterial.FragmentShader, "fragment shader\n")
}

func Test_LoadMaterial_IgnoresFileLoadIfShaderSourceExists(t *testing.T) {
	render.DefaultMaterial.VertexShader = "default vertex shader"
	render.DefaultMaterial.FragmentShader = "default fragment shader"

	getTestGraphical()

	// Check shader code of default material, loaded by NewGraphical
	assert.Equal(t, render.DefaultMaterial.VertexShader, "default vertex shader")
	assert.Equal(t, render.DefaultMaterial.FragmentShader, "default fragment shader")
}

func Test_LoadMaterial_FallsBackToDefaultMeshIfNoShaderFilesFound(t *testing.T) {
	render.DefaultMaterial.VertexShader = "default vertex shader"
	render.DefaultMaterial.FragmentShader = "default fragment shader"

	material := &render.Material{
		Name:    "testMat",
		Shaders: "missing",
	}

	graphical, _, _ := getTestGraphical()
	graphical.LoadMaterial(material)

	assert.Equal(t, material.Name, render.DefaultMaterial.Name)
	assert.Equal(t, render.DefaultMaterial.VertexShader, material.VertexShader)
	assert.Equal(t, render.DefaultMaterial.FragmentShader, material.FragmentShader)
}

func Test_SetUpEntity_SetsDefaultMeshIfNoneSpecified(t *testing.T) {
	graphical, _, _ := getTestGraphical()

	entity := core.NewEntity()
	entity.AddComponent(new(components.Visual))

	graphical.SetUpEntity(entity)

	visual := components.GetVisual(entity)
	assert.Equal(t, render.DefaultMesh.Name, visual.MeshName)
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

	visual := components.GetVisual(entity)
	assert.Equal(t, render.DefaultMaterial.Name, visual.MaterialName)
}

func Test_SetUpEntity_LinksLoadedMaterialIfNameMatches(t *testing.T) {
}

func Test_SetUpEntity_TellsRendererToLoadNewMaterial(t *testing.T) {
}

func Test_Update_ConfiguresQueueWithViewData(t *testing.T) {
	graphical, renderer, _ := getTestGraphical()

	camera := core.NewCamera()
	camera.Position = math3d.Vector{5, 5, 5}
	camera.LookAt = math3d.Vector{1, 2, 3}

	graphical.Update(camera, 0)

	assert.NotNil(t, renderer.queueRendered)

	assert.Equal(t, camera.ProjectionMatrix(), renderer.queueRendered.ProjectionMatrix)
	assert.Equal(t, camera.ViewMatrix(), renderer.queueRendered.ViewMatrix)
}

func Test_Update_RendersAllVisualEntityMeshesWithMaterials(t *testing.T) {
	graphical, renderer, entityDb := getTestGraphical()

	entity := core.NewEntity()
	entity.AddComponent(new(components.Visual))
	entityDb.RegisterEntity(entity)

	graphical.Update(core.NewCamera(), 0)

	assert.NotNil(t, renderer.queueRendered)

	assert.Equal(t, 1, len(renderer.queueRendered.RenderOperations()))

	renderOp := renderer.queueRendered.RenderOperations()[0]

	assert.Equal(t, render.DefaultMesh, renderOp.Mesh)
	assert.Equal(t, render.DefaultMaterial, renderOp.Material)
	assert.Equal(t, math3d.IdentityMatrix(), renderOp.Transform)
}
