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

type RenderState struct {
	Projection, View math3d.Matrix
}

func (self *OpenGLRenderer) Render(renderQueue *render.RenderQueue) {
	renderState := RenderState{
		Projection: renderQueue.ProjectionMatrix,
		View:       renderQueue.ViewMatrix,
	}

	for _, renderOp := range renderQueue.RenderOperations() {
		self.renderOne(renderOp, renderState)
	}
}

func (self *OpenGLRenderer) renderOne(operation render.RenderOperation, renderState RenderState) {
	mesh := operation.Mesh
	material := operation.Material
	transform := operation.Transform

	vertexArrayObj := mesh.VertexArrayObj.(gl.VertexArray)
	vertexArrayObj.Bind()

	material.ShaderProgram.SetUniformMatrix(
		"modelViewProjection",
		renderState.Projection.Times(renderState.View).Times(transform),
	)
	material.ShaderProgram.Use()

	gl.DrawArrays(gl.TRIANGLES, 0, len(mesh.VertexList)*3)
}

func (self *OpenGLRenderer) FinishRender() {
}
