package platform

import (
	"core"
	"github.com/go-gl/gl"
	"math3d"
)

type OpenGLRenderer struct {
	// Implements the core.Renderer interface
}

func (self *OpenGLRenderer) LoadMesh(mesh *core.Mesh) {
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

func (self *OpenGLRenderer) LoadMaterial(material *core.Material) {
	material.ShaderProgram = NewGLSLProgram(material.VertexShader, material.FragmentShader)
}

func (self *OpenGLRenderer) BeginRender() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func (self *OpenGLRenderer) Render(mesh *core.Mesh, material *core.Material) {
	vertexArrayObj := mesh.VertexArrayObj.(gl.VertexArray)
	vertexArrayObj.Bind()

	// The projection and view need to be set once when rendering,
	// only the model matrix changes per entity rendered.
	projection := math3d.Perspective(45.0, 4.0/3.0, 0.1, 100.0)
	view := math3d.LookAt(
		4, 3, 3,
		0, 0, 0,
		0, 1, 0,
	)
	model := math3d.IdentityMatrix()

	material.ShaderProgram.SetUniformMatrix("modelViewProjection", projection.Times(view).Times(model))
	material.ShaderProgram.Use()

	gl.DrawArrays(gl.TRIANGLES, 0, len(mesh.VertexList) * 3)
}

func (self *OpenGLRenderer) FinishRender() {
}
