package platform

import (
	"github.com/go-gl/gl"
	"log"
	"math3d"
)

type uniformMap map[string]gl.UniformLocation

type GLSLProgram struct {
	// Implements core.GPUProgram
	program        gl.Program
	vertexShader   gl.Shader
	fragmentShader gl.Shader

	// Keep track of the known uniforms for this shader
	// so we aren't constantly doing a GPU query
	uniforms uniformMap
}

func NewGLSLProgram(vertexProgram, fragmentProgram string) *GLSLProgram {
	program := new(GLSLProgram)
	program.uniforms = make(uniformMap)

	program.program = gl.CreateProgram()
	program.vertexShader = program.buildShader(gl.VERTEX_SHADER, vertexProgram)
	program.fragmentShader = program.buildShader(gl.FRAGMENT_SHADER, fragmentProgram)

	program.attachAndLinkShaders()
	program.validateShaders()

	return program
}

func (self *GLSLProgram) Use() {
	self.program.Use()
}

func (self *GLSLProgram) SetUniformMatrix(uniformName string, matrix math3d.Matrix) {
	uniformLoc := self.getUniformLocation(uniformName)

	if uniformLoc > -1 {
		uniformLoc.UniformMatrix4fv(false, matrix)
	}
}

func (self *GLSLProgram) getUniformLocation(uniformName string) gl.UniformLocation {
	var loc gl.UniformLocation

	if loc, ok := self.uniforms[uniformName]; !ok {
		loc = self.program.GetUniformLocation(uniformName)
		self.uniforms[uniformName] = loc
	}

	return loc
}

func (self *GLSLProgram) attachAndLinkShaders() {
	self.program.AttachShader(self.vertexShader)
	self.program.AttachShader(self.fragmentShader)
	self.program.Link()

	linkStatus := self.program.Get(gl.LINK_STATUS)
	if linkStatus != 1 {
		log.Print("Failed to link programs, status: ", linkStatus)
		log.Panic("Info log: ", self.program.GetInfoLog())
	}
}

func (self *GLSLProgram) validateShaders() {
	self.program.Validate()
	validationStatus := self.program.Get(gl.VALIDATE_STATUS)
	if validationStatus != 1 {
		log.Panic("Program validation failed: ", validationStatus)
	}
}

func (self *GLSLProgram) buildShader(shaderType gl.GLenum, source string) gl.Shader {
	shader := gl.CreateShader(shaderType)
	shader.Source(source)
	shader.Compile()

	compileStatus := shader.Get(gl.COMPILE_STATUS)
	if compileStatus != 1 {
		log.Print("Shader compilation status: ", compileStatus)
		log.Print("Info Log: ", shader.GetInfoLog())
		log.Panic("Error creating shader!")
	}

	return shader
}