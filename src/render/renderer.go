package render

// Any renderer must adhere to this interface.
type Renderer interface {

	// LoadMesh takes the given Mesh and loads it such that it's now renderable
	LoadMesh(mesh *Mesh)

	// LoadMaterial takes a Material definition and sets up appropriate hardware values
	LoadMaterial(material *Material)

	//
	// Render path
	//
	BeginRender()
	Render(renderQueue *RenderQueue)
	FinishRender()
}
