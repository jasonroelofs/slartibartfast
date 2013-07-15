package platform

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glu"
	"log"
	"math3d"
	"render"
)

type OpenGLRenderer struct {
	// Implements the render.Renderer interface
}

func (self *OpenGLRenderer) LoadMesh(mesh *render.Mesh) {
	if len(mesh.VertexList) == 0 {
		log.Println("WARNING Stopping load of mesh [", mesh.Name, "] with no verticies")
		return
	}

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

	if len(mesh.IndexList) > 0 {
		indexBuffer := gl.GenBuffer()
		indexBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mesh.IndexList)*4, mesh.IndexList, gl.STATIC_DRAW)

		mesh.IndexBuffer = indexBuffer
	}

	mesh.VertexArrayObj = vertexArrayObj
	mesh.VertexBuffer = vertexBuffer
}

func (self *OpenGLRenderer) LoadMaterial(material *render.Material) {
	material.Shader.Program = NewGLSLProgram(material.Shader.Vertex, material.Shader.Fragment)

	if material.IsCubeMap {
		material.Texture.Id = self.loadCubeMap(material)
	} else if material.Texture != nil {
		material.Texture.Id = self.LoadTexture(material.Texture)
	}

	self.checkErrors()
}

func (self *OpenGLRenderer) loadCubeMap(material *render.Material) gl.Texture {
	glTexture := gl.GenTexture()
	glTexture.Bind(gl.TEXTURE_CUBE_MAP)
	defer glTexture.Unbind(gl.TEXTURE_CUBE_MAP)

	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	self.glTexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_X, material.CubeMap[render.CubeFace_Right])
	self.glTexImage2D(gl.TEXTURE_CUBE_MAP_NEGATIVE_X, material.CubeMap[render.CubeFace_Left])

	self.glTexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_Y, material.CubeMap[render.CubeFace_Top])
	self.glTexImage2D(gl.TEXTURE_CUBE_MAP_NEGATIVE_Y, material.CubeMap[render.CubeFace_Bottom])

	self.glTexImage2D(gl.TEXTURE_CUBE_MAP_POSITIVE_Z, material.CubeMap[render.CubeFace_Front])
	self.glTexImage2D(gl.TEXTURE_CUBE_MAP_NEGATIVE_Z, material.CubeMap[render.CubeFace_Back])

	return glTexture
}

func (self *OpenGLRenderer) LoadTexture(texture *render.Texture) gl.Texture {
	glTexture := gl.GenTexture()
	glTexture.Bind(gl.TEXTURE_2D)
	defer glTexture.Unbind(gl.TEXTURE_2D)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)

	self.glTexImage2D(gl.TEXTURE_2D, texture)

	gl.GenerateMipmap(gl.TEXTURE_2D)

	return glTexture
}

func (self *OpenGLRenderer) glTexImage2D(textureType gl.GLenum, texture *render.Texture) {
	gl.TexImage2D(textureType, 0, gl.RGB,
		texture.Image.Width(), texture.Image.Height(), 0,
		gl.BGR, gl.UNSIGNED_BYTE,
		texture.Image.Bytes(),
	)
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

	// No attributes? no loaded or empty? Better way to handle this than spamming
	// the console?
	if mesh.VertexArrayObj == nil {
		log.Println("WARNING: Trying to render an invalid mesh", mesh)
		return
	}

	vertexArrayObj := mesh.VertexArrayObj.(gl.VertexArray)
	vertexArrayObj.Bind()

	material.Shader.Program.Use()
	material.Shader.Program.SetUniformMatrix(
		"modelViewProjection",
		renderState.Projection.Times(renderState.View).Times(transform),
	)

	if material.Texture != nil {
		glTexture := material.Texture.Id.(gl.Texture)
		gl.ActiveTexture(gl.TEXTURE0)

		if !material.IsCubeMap {
			glTexture.Bind(gl.TEXTURE_2D)
			defer glTexture.Unbind(gl.TEXTURE_2D)
			material.Shader.Program.SetUniformUnit("textureSampler", 0)
		} else {
			gl.Disable(gl.DEPTH_TEST)
			defer gl.Enable(gl.DEPTH_TEST)

			glTexture.Bind(gl.TEXTURE_CUBE_MAP)
			defer glTexture.Unbind(gl.TEXTURE_CUBE_MAP)
			material.Shader.Program.SetUniformUnit("cubeMap", 0)
		}
	}

	if len(mesh.IndexList) == 0 {
		gl.DrawArrays(gl.TRIANGLES, 0, len(mesh.VertexList)*3)
	} else {
		gl.DrawElements(gl.TRIANGLES, len(mesh.IndexList), gl.UNSIGNED_INT, nil)
	}
}

func (self *OpenGLRenderer) FinishRender() {
	gl.Disable(gl.DEPTH_TEST)
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
