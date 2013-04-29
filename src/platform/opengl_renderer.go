package platform

import (
	"core"
	"github.com/go-gl/gl"
)

type OpenGLRenderer struct {
	// Implements the core.Renderer interface
}

func (self *OpenGLRenderer) LoadMesh(mesh *core.Mesh) {
	vertexArrayObj := gl.GenVertexArray()
	vertexArrayObj.Bind()

	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	defer vertexBuffer.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.VertexList) * 4, mesh.VertexList, gl.STATIC_DRAW)

	colorBuffer := gl.GenBuffer()
	colorBuffer.Bind(gl.ARRAY_BUFFER)
	defer colorBuffer.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.ColorList) * 4, mesh.ColorList, gl.STATIC_DRAW)

	//mesh.VertexArrayObj = vertexArrayObj.(gl.GLuint)
	//mesh.VertexBuffer = vertexBuffer.(gl.GLuint)
	//mesh.ColorBuffer = colorBuffer.(gl.GLuint)
}
