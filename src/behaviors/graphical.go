package behaviors

import (
	"components"
	"core"
	"io/ioutil"
	"log"
	"render"
)

type MeshMap map[string]*render.Mesh
type MaterialMap map[string]*render.Material

type Graphical struct {
	renderer  render.Renderer
	entitySet *core.EntitySet
	meshes    MeshMap
	materials MaterialMap
}

func NewGraphical(renderer render.Renderer, entityDB *core.EntityDB) *Graphical {
	obj := Graphical{}
	obj.renderer = renderer
	obj.entitySet = entityDB.RegisterListener(&obj, components.TRANSFORM, components.VISUAL)

	obj.meshes = make(MeshMap)
	obj.materials = make(MaterialMap)

	obj.LoadMesh(render.DefaultMesh)
	obj.LoadMaterial(render.DefaultMaterial)

	return &obj
}

// EntityListener
func (self *Graphical) SetUpEntity(entity *core.Entity) {
	visual := components.GetVisual(entity)
	self.linkMeshToVisual(visual)
	self.linkMaterialToVisual(visual)
}

// Mesh loading
func (self *Graphical) linkMeshToVisual(visual *components.Visual) {
	if visual.MeshName == "" {
		visual.MeshName = render.DefaultMesh.Name
	} else {
		// find existing mesh?
		// or load mesh by name
	}
}

// Material loading
func (self *Graphical) linkMaterialToVisual(visual *components.Visual) {
	if visual.MaterialName == "" {
		visual.MaterialName = render.DefaultMaterial.Name
	} else {
		// Nothing or Load new Material
	}
}

// LoadMesh takes a given Mesh object and ensures it's contents are
// loaded into the renderer for future use.
func (self *Graphical) LoadMesh(mesh *render.Mesh) {
	log.Println("Loading Mesh", mesh.Name)
	self.renderer.LoadMesh(mesh)
	self.meshes[mesh.Name] = mesh
}

// LoadMaterial takes a Material object and ensures it is available to the
// renderer for future use.
func (self *Graphical) LoadMaterial(material *render.Material) {
	log.Println("Loading Material", material.Name)
	self.loadShadersIntoMaterial(material)
	self.renderer.LoadMaterial(material)
	self.materials[material.Name] = material
}

func (self *Graphical) loadShadersIntoMaterial(material *render.Material) {
	if material.VertexShader != "" && material.FragmentShader != "" {
		return
	}

	baseShaderPath := defaults.LoadPath + "/shaders/" + material.Shaders

	vertPath := baseShaderPath + ".vert"
	log.Println("Loading vertex shader", vertPath)

	// This error handling isn't quite right. If vertex succeeds but fragment fails we have
	// two shaders that don't work together. Maybe falling back to default isn't a good idea
	// and should just error out here?

	vertSource, err := ioutil.ReadFile(vertPath)
	if err != nil {
		log.Println("Error loading vertex shader", vertPath, err, "Reverting to default Material")
		material.Name = render.DefaultMaterial.Name
		vertSource = []byte(render.DefaultMaterial.VertexShader)
	}

	fragPath := baseShaderPath + ".frag"
	log.Println("Loading fragment shader", fragPath)

	fragSource, err := ioutil.ReadFile(fragPath)
	if err != nil {
		log.Println("Error loading fragment shader", fragPath, err, "Reverting to default Material")
		material.Name = render.DefaultMaterial.Name
		fragSource = []byte(render.DefaultMaterial.FragmentShader)
	}

	material.VertexShader = string(vertSource)
	material.FragmentShader = string(fragSource)
}

// Update is called every Game tick
func (self *Graphical) Update(camera *core.Camera, deltaT float32) {
	self.renderer.BeginRender()
	var visual *components.Visual

	renderQueue := render.NewRenderQueue()

	renderQueue.ProjectionMatrix = camera.ProjectionMatrix()
	renderQueue.ViewMatrix = camera.ViewMatrix()

	for _, entity := range self.entitySet.Entities {
		visual = components.GetVisual(entity)

		renderQueue.Add(render.RenderOperation{
			Mesh:      self.meshes[visual.MeshName],
			Material:  self.materials[visual.MaterialName],
			Transform: components.GetTransform(entity).TransformMatrix(),
		})
	}

	self.renderer.Render(renderQueue)

	self.renderer.FinishRender()
}
