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

	// Load some test shaders
	vertexProgram := `
#version 150

in vec3 vertexPosition;
in vec3 in_color;
out vec3 vert_color;

void main() {
	gl_Position.xyz = vertexPosition;
	gl_Position.w = 1.0;
	vert_color = in_color;
}
`

	fragmentProgram := `
#version 150

in vec3 vert_color;
out vec4 frag_color;

void main() {
	frag_color = vec4(vert_color, 0);
}
`
	program := NewGLSLProgram(vertexProgram, fragmentProgram)
	program.Use()

	mesh.VertexArrayObj = vertexArrayObj
	mesh.VertexBuffer = vertexBuffer
}

func (self *OpenGLRenderer) BeginRender() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Set the camera manually for now
	// A few units away and looking at the origin
	//	glu.LookAt(
	//		0, 0, -10,
	//		0, 0, 0,
	//		0, 1, 0,
	//	)
}

func (self *OpenGLRenderer) Render(mesh core.Mesh) {
	vertexArrayObj := mesh.VertexArrayObj.(gl.VertexArray)
	vertexArrayObj.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func (self *OpenGLRenderer) FinishRender() {
}
