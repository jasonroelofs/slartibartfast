package render

import (
	"io/ioutil"
	"log"
)

type shaderMap map[string]*Shader

type ShaderLoader struct {
	loadedShaders shaderMap
}

func NewShaderLoader() *ShaderLoader {
	loader := new(ShaderLoader)
	loader.loadedShaders = make(shaderMap)
	return loader
}

func (self *ShaderLoader) Load(shaderName string) *Shader {
	shader := Shader{
		Vertex:   self.loadProgram(shaderName + ".vert"),
		Fragment: self.loadProgram(shaderName + ".frag"),
	}

	self.loadedShaders[shaderName] = &shader

	return &shader
}

func (self *ShaderLoader) loadProgram(programName string) string {
	log.Println("Loading shader program", programName)

	path := Defaults.LoadPath + "/shaders/" + programName

	source, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Error loading shader", path, "", err)
	}

	return string(source)
}
