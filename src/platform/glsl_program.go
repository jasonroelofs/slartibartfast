package platform

import (
	"github.com/go-gl/gl"
	"log"
)

type GLSLProgram struct {
	program        gl.Program
	vertexShader   gl.Shader
	fragmentShader gl.Shader
}

func NewGLSLProgram(vertexProgram, fragmentProgram string) *GLSLProgram {
	program := new(GLSLProgram)
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

func (self *GLSLProgram) GetUniformLocation(uniform string) gl.UniformLocation {
	return self.program.GetUniformLocation(uniform)
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
