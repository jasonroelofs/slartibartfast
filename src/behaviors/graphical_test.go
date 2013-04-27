package behaviors

import (
	"components"
	"core"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

type TestRenderer struct {
	loadedMesh *core.Mesh
}

func (self *TestRenderer) LoadMesh(mesh *core.Mesh) {
	self.loadedMesh = mesh
}

func getTestGraphical() (*Graphical, *TestRenderer, *core.EntityDB) {
	entityDB := new(core.EntityDB)
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

	assert.Equal(t, core.DefaultMesh, renderer.loadedMesh)
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
