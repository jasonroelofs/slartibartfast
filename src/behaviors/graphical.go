package behaviors

import (
	"components"
	"core"
	"log"
	"render"
)

type MeshMap map[string]*render.Mesh
type MaterialMap map[string]*render.Material

type Graphical struct {
	renderer  render.Renderer
	entitySet *core.EntitySet

	materialLoader *render.MaterialLoader

	meshes MeshMap
}

func NewGraphical(renderer render.Renderer, entityDB *core.EntityDB) *Graphical {
	obj := Graphical{}
	obj.renderer = renderer
	obj.entitySet = entityDB.RegisterListener(&obj, components.TRANSFORM, components.VISUAL)

	obj.materialLoader = render.NewMaterialLoader()

	obj.meshes = make(MeshMap)

	obj.LoadMesh(render.DefaultMesh)
	obj.LoadMaterial(render.DefaultMaterial)

	return &obj
}

// SetUpEntity :: EntityListener. Will load up the mesh hooked to the given entity's
// Visual component if one was given directly instead of via MeshName
func (self *Graphical) SetUpEntity(entity *core.Entity) {
	visual := components.GetVisual(entity)

	self.checkDefaultsOnVisual(visual)

	if visual.Mesh != nil {
		self.LoadMesh(visual.Mesh)
	}
}

// TearDownEntity :: EntityListener. Cleans up any visual resources that are
// directly linked to this Entity's Visual component.
func (self *Graphical) TearDownEntity(entity *core.Entity) {
}

func (self *Graphical) checkDefaultsOnVisual(component *components.Visual) {
	// Ensure we load something for this Visual.
	if component.Mesh == nil && component.MeshName == "" {
		component.MeshName = render.DefaultMeshName
	}

	if component.MaterialName == "" {
		component.MaterialName = render.DefaultMaterialName
	}
}

// LoadMesh takes a given Mesh object and ensures it's contents are
// loaded into the renderer for future use.
func (self *Graphical) LoadMesh(mesh *render.Mesh) {
	log.Println("Loading Mesh", mesh.Name)

	self.renderer.LoadMesh(mesh)

	if mesh.Name != "" {
		self.meshes[mesh.Name] = mesh
	}
}

// LoadMaterial takes a MaterialDef object and ensures it is available to the
// renderer for future use.
func (self *Graphical) LoadMaterial(materialDef render.MaterialDef) {
	material := self.materialLoader.Load(materialDef)
	self.renderer.LoadMaterial(material)
}

// Update is called every Game tick
func (self *Graphical) Update(camera *core.Camera, deltaT float32) {
	self.renderer.BeginRender()
	var visual *components.Visual

	renderQueue := render.NewRenderQueue()

	renderQueue.ProjectionMatrix = camera.ProjectionMatrix()
	renderQueue.ViewMatrix = camera.ViewMatrix()

	for _, entity := range self.entitySet.Entities() {
		visual = components.GetVisual(entity)

		renderQueue.Add(render.RenderOperation{
			Mesh:      self.meshFromVisual(visual),
			Material:  self.materialLoader.Get(visual.MaterialName),
			Transform: components.GetTransform(entity).TransformMatrix(),
		})
	}

	self.renderer.Render(renderQueue)

	self.renderer.FinishRender()
}

// Either use the mesh that's directly linked to the Visual component or look
// up preloaded meshes by name
func (self *Graphical) meshFromVisual(visual *components.Visual) *render.Mesh {
	if visual.Mesh != nil {
		return visual.Mesh
	} else {
		return self.meshes[visual.MeshName]
	}
}
