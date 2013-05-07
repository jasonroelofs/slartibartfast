package behaviors

import (
	"components"
	"core"
	"io/ioutil"
	"log"
)

type MeshMap map[string]*core.Mesh
type MaterialMap map[string]*core.Material

type Graphical struct {
	renderer  core.Renderer
	entitySet *core.EntitySet
	meshes    MeshMap
	materials MaterialMap
}

func NewGraphical(renderer core.Renderer, entityDB *core.EntityDB) *Graphical {
	obj := Graphical{}
	obj.renderer = renderer
	obj.entitySet = entityDB.RegisterListener(&obj, components.TRANSFORM, components.VISUAL)

	obj.meshes = make(MeshMap)
	obj.materials = make(MaterialMap)

	obj.LoadMesh(core.DefaultMesh)
	obj.LoadMaterial(core.DefaultMaterial)

	return &obj
}

// EntityListener
func (self *Graphical) SetUpEntity(entity *core.Entity) {
	visual := entity.GetComponent(components.VISUAL).(*components.Visual)
	self.linkMeshToVisual(visual)
	self.linkMaterialToVisual(visual)
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

// Material loading
func (self *Graphical) linkMaterialToVisual(visual *components.Visual) {
	if visual.MaterialName == "" {
		visual.MaterialName = core.DefaultMaterial.Name
	} else {
		// Nothing or Load new Material
	}
}

// LoadMesh takes a given Mesh object and ensures it's contents are
// loaded into the renderer for future use.
func (self *Graphical) LoadMesh(mesh *core.Mesh) {
	log.Println("Loading Mesh", mesh.Name)
	self.renderer.LoadMesh(mesh)
	self.meshes[mesh.Name] = mesh
}

// LoadMaterial takes a Material object and ensures it is available to the
// renderer for future use.
func (self *Graphical) LoadMaterial(material *core.Material) {
	log.Println("Loading Material", material.Name)
	self.loadShadersIntoMaterial(material)
	self.renderer.LoadMaterial(material)
	self.materials[material.Name] = material
}

func (self *Graphical) loadShadersIntoMaterial(material *core.Material) {
	if material.VertexShader != "" && material.FragmentShader != "" {
		return
	}

	baseShaderPath := defaults.LoadPath + "/shaders/" + material.Shaders

	vertPath := baseShaderPath + ".vert"
	log.Println("Loading vertex shader", vertPath)

	vertSource, err := ioutil.ReadFile(vertPath)
	if err != nil {
		log.Println("Error loading vertex shader", vertPath, err, "Reverting to default Material")
		material.Name = core.DefaultMaterial.Name
		vertSource = []byte(core.DefaultMaterial.VertexShader)
	}

	fragPath := baseShaderPath + ".frag"
	log.Println("Loading fragment shader", fragPath)

	fragSource, err := ioutil.ReadFile(fragPath)
	if err != nil {
		log.Println("Error loading fragment shader", fragPath, err, "Reverting to default Material")
		material.Name = core.DefaultMaterial.Name
		fragSource = []byte(core.DefaultMaterial.FragmentShader)
	}

	material.VertexShader = string(vertSource)
	material.FragmentShader = string(fragSource)
}

// Update is called every Game tick
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
		self.renderer.Render(self.meshes[visual.MeshName], self.materials[visual.MaterialName])
	}

	self.renderer.FinishRender()
}
