package platform

import (
	"github.com/go-gl/gl"
	"math3d"
	"render"
)

type OpenGLRenderer struct {
	// Implements the render.Renderer interface
}

func (self *OpenGLRenderer) LoadMesh(mesh *render.Mesh) {
	vertexArrayObj := gl.GenVertexArray()
	vertexArrayObj.Bind()

	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.VertexList)*4, mesh.VertexList, gl.STATIC_DRAW)

	attribLoc := gl.AttribLocation(0)
	attribLoc.EnableArray()
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

	if len(mesh.ColorList) > 0 {
		colorBuffer := gl.GenBuffer()
		colorBuffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, len(mesh.ColorList)*4, mesh.ColorList, gl.STATIC_DRAW)

		attribLoc := gl.AttribLocation(1)
		attribLoc.EnableArray()
		attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

		mesh.ColorBuffer = colorBuffer
	}

	mesh.VertexArrayObj = vertexArrayObj
	mesh.VertexBuffer = vertexBuffer
}

func (self *OpenGLRenderer) LoadMaterial(material *render.Material) {
	material.ShaderProgram = NewGLSLProgram(material.VertexShader, material.FragmentShader)
}

func (self *OpenGLRenderer) BeginRender() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func (self *OpenGLRenderer) Render(renderQueue *render.RenderQueue) {
	for _, renderOp := range renderQueue.RenderOperations() {
		self.renderOne(renderOp)
	}
}

func (self *OpenGLRenderer) renderOne(operation render.RenderOperation) {
	mesh := operation.Mesh
	material := operation.Material
	transform := operation.Transform

	vertexArrayObj := mesh.VertexArrayObj.(gl.VertexArray)
	vertexArrayObj.Bind()

	// The projection and view need to be set once when rendering,
	// only the model matrix changes per entity rendered.
	projection := math3d.Perspective(45.0, 4.0/3.0, 0.1, 100.0)
	view := math3d.LookAt(
		math3d.Vector{14, 13, 13},
		math3d.Vector{0, 0, 0},
		math3d.Vector{0, 1, 0},
	)

	material.ShaderProgram.SetUniformMatrix(
		"modelViewProjection", projection.Times(view).Times(transform))
	material.ShaderProgram.Use()

	gl.DrawArrays(gl.TRIANGLES, 0, len(mesh.VertexList)*3)
}

func (self *OpenGLRenderer) FinishRender() {
}
