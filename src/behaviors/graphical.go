package behaviors

import (
	"components"
	"core"
)

type MeshMap map[string]core.Mesh

type Graphical struct {
	renderer  core.Renderer
	entitySet *core.EntitySet
	meshes    MeshMap
}

func NewGraphical(renderer core.Renderer, entityDB *core.EntityDB) *Graphical {
	obj := Graphical{}
	obj.renderer = renderer
	obj.entitySet = entityDB.RegisterListener(&obj, components.TRANSFORM, components.VISUAL)
	obj.meshes = make(MeshMap)
	obj.loadMesh(core.DefaultMesh)

	return &obj
}

// EntityListener
func (self *Graphical) SetUpEntity(entity *core.Entity) {
	visual := entity.GetComponent(components.VISUAL).(*components.Visual)
	self.linkMeshToVisual(visual)
}

// Mesh loading
func (self *Graphical) linkMeshToVisual(visual *components.Visual) {
	if visual.MeshName == "" {
		visual.MeshName = core.DefaultMesh.Name
	} else {
		// find existing mesh?
		// or load mesh by name
	}
}

func (self *Graphical) loadMesh(mesh *core.Mesh) {
	self.renderer.LoadMesh(mesh)
	self.meshes[mesh.Name] = *mesh
}

// Game tick
func (self *Graphical) Update(deltaT float64) {
	self.renderer.BeginRender()
	var visual *components.Visual

	// Render the Visible's vertex / index arrays for each entity
	// For each entity in entitySet
	// - Build a render operation for that entity's data
	// Tell renderer to render the set of render ops

	// For now (prototyping) just render the mesh

	for _, entity := range self.entitySet.Entities {
		visual = entity.GetComponent(components.VISUAL).(*components.Visual)
		self.renderer.Render(self.meshes[visual.MeshName])
	}

	self.renderer.FinishRender()
}
