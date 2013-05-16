package platform

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glu"
	"math3d"
	"render"
	"log"
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

	if len(mesh.UVList) > 0 {
		uvBuffer := gl.GenBuffer()
		uvBuffer.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, len(mesh.UVList)*4, mesh.UVList, gl.STATIC_DRAW)

		attribLoc := gl.AttribLocation(1)
		attribLoc.EnableArray()
		attribLoc.AttribPointer(2, gl.FLOAT, false, 0, nil)

		mesh.UVBuffer = uvBuffer
	} else if len(mesh.ColorList) > 0 {
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
	material.Shader.Program = NewGLSLProgram(material.Shader.Vertex, material.Shader.Fragment)
	if material.Texture != nil {
		material.Texture.Id = self.LoadTexture(material.Texture)
	}
}

func (self *OpenGLRenderer) LoadTexture(texture *render.Texture) gl.Texture {
	glTexture := gl.GenTexture()
	glTexture.Bind(gl.TEXTURE_2D)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB,
		texture.Image.Width(), texture.Image.Height(), 0, gl.RGB, gl.UNSIGNED_BYTE,
		texture.Image.Bytes())

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	self.checkErrors()

	return glTexture
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

	material.Shader.Program.SetUniformMatrix(
		"modelViewProjection",
		renderState.Projection.Times(renderState.View).Times(transform),
	)
	material.Shader.Program.Use()

	if material.Texture != nil {
		glTexture := material.Texture.Id.(gl.Texture)
		gl.ActiveTexture(gl.TEXTURE0)
		glTexture.Bind(gl.TEXTURE_2D)
		defer glTexture.Unbind(gl.TEXTURE_2D)

		material.Shader.Program.SetUniformUnit("textureSampler", 0)
	}

	gl.DrawArrays(gl.TRIANGLES, 0, len(mesh.VertexList)*3)
}

func (self *OpenGLRenderer) FinishRender() {
}

func (self *OpenGLRenderer) checkErrors() {
	e := gl.GetError()
	for e != gl.NO_ERROR {
		errString, err := glu.ErrorString(e)
		if err != nil {
			log.Println("Invalid error code found", err)
		} else {
			log.Printf("GLError %v => %s", e, errString)
		}

		e = gl.GetError()
	}
}
