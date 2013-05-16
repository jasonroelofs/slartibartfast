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

	SetUniformMatrix(uniformName string, matrix math3d.Matrix)
	SetUniformUnit(uniformName string, unitIndex int)
}
