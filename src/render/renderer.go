package render

import (
	"math3d"
)

// All the interfaces required for a renderer subsystem to be compliant in this system

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
	Render(mesh *Mesh, material *Material)
	FinishRender()
}

// GPU Shader Program
type GPUProgram interface {
	// Use tells the underlying render system to use this program
	Use()

	// SetUniformMatrix sets the uniform value in the current shaders to the given Matrix
	SetUniformMatrix(uniformName string, matrix math3d.Matrix)
}
