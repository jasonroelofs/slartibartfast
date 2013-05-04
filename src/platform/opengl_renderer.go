package platform

import (
	"core"
	"github.com/go-gl/gl"
	"math3d"
)

type OpenGLRenderer struct {
	// Implements the core.Renderer interface
}

var mvpLocation gl.UniformLocation

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

uniform mat4 MVP;

void main() {
	gl_Position = MVP * vec4(vertexPosition, 1.0);
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

	mvpLocation = program.GetUniformLocation("MVP")

	mesh.VertexArrayObj = vertexArrayObj
	mesh.VertexBuffer = vertexBuffer
}

func (self *OpenGLRenderer) LoadMaterial(material *core.Material) {
}

func (self *OpenGLRenderer) BeginRender() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func (self *OpenGLRenderer) Render(mesh core.Mesh) {
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

	mvp := projection.Times(view).Times(model)

	mvpLocation.UniformMatrix4fv(false, mvp)

	gl.DrawArrays(gl.TRIANGLES, 0, len(mesh.VertexList) * 3)
}

func (self *OpenGLRenderer) FinishRender() {
}
