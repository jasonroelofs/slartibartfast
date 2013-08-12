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
	unloadedMesh   *render.Mesh
	loadedMaterial *render.Material
	queueRendered  *render.RenderQueue
}

func (self *TestRenderer) LoadMesh(mesh *render.Mesh) {
	self.loadedMesh = mesh
}

func (self *TestRenderer) UnloadMesh(mesh *render.Mesh) {
	self.unloadedMesh = mesh
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
	entityDB := core.NewEntityDB()
	renderer := new(TestRenderer)

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

	assert.Equal(t, render.DefaultMaterial.Name, renderer.loadedMaterial.Name)
}

func Test_SetUpEntity_TellsRendererToLoadNewMeshFromVisual(t *testing.T) {
	graphical, renderer, entityDb := getTestGraphical()

	mesh := &render.Mesh{}
	entity := core.NewEntity()
	entity.AddComponent(&components.Visual{Mesh: mesh})
	entityDb.RegisterEntity(entity)

	assert.Equal(t, mesh, renderer.loadedMesh)
	assert.NotEqual(t, mesh, graphical.meshes[mesh.Name])
}

func Test_SetUpEntity_TellsRendererToLoadNewMaterial(t *testing.T) {
}

func Test_TearDownEntity_TellsRendererToUnload(t *testing.T) {
	_, renderer, entityDb := getTestGraphical()

	mesh := &render.Mesh{}
	entity := core.NewEntity()
	entity.AddComponent(&components.Visual{Mesh: mesh})
	entityDb.RegisterEntity(entity)

	entity.Destroy()

	assert.Equal(t, mesh, renderer.unloadedMesh)
}

func Test_Update_ConfiguresQueueWithViewData(t *testing.T) {
	graphical, renderer, _ := getTestGraphical()

	camera := core.NewCamera()
	camera.SetPosition(math3d.Vector{5, 5, 5})
	camera.LookAt(math3d.Vector{1, 2, 3})

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
	assert.Equal(t, render.DefaultMaterial.Name, renderOp.Material.Name)
	assert.Equal(t, math3d.IdentityMatrix(), renderOp.Transform)
}

func Test_Update_UsesMeshLinkedToVisualIfExists(t *testing.T) {
	graphical, renderer, entityDb := getTestGraphical()

	mesh := &render.Mesh{}

	entity := core.NewEntity()
	entity.AddComponent(&components.Visual{Mesh: mesh})
	entityDb.RegisterEntity(entity)

	graphical.Update(core.NewCamera(), 0)
	renderOp := renderer.queueRendered.RenderOperations()[0]

	assert.Equal(t, mesh, renderOp.Mesh)
}
