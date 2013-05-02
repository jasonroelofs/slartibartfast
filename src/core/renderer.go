package core

// Any renderer must adhere to this interface.
type Renderer interface {

	// LoadMesh takes the given Mesh and loads it such that it's now renderable
	LoadMesh(mesh *Mesh)

	//
	// Render path
	//
	BeginRender()
	Render(mesh Mesh)
	FinishRender()
}
