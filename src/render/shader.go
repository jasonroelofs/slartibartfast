package render

import "math3d"

type Shader struct {
	// Source of the appropriate shaders
	Vertex   string
	Fragment string

	// Link to the program loaded into the GPU
	Program  GPUProgram
}

type GPUProgram interface {
	// Use tells the underlying render system to use this program
	Use()

	// SetUniformMatrix sets the uniform value in the current shaders to the given Matrix
	SetUniformMatrix(uniformName string, matrix math3d.Matrix)
}
