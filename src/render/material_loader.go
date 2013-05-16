package render

import (
	"log"
)

const (
	defaultMaterialName = ""
)

type materialMap map[string]*Material

type MaterialLoader struct {
	shaderLoader    *ShaderLoader
	textureLoader   *TextureLoader

	loadedMaterials materialMap
}

func NewMaterialLoader() *MaterialLoader {
	loader := new(MaterialLoader)
	loader.loadedMaterials = make(materialMap)

	loader.shaderLoader = NewShaderLoader()
	loader.textureLoader = NewTextureLoader()

	return loader
}

func (self *MaterialLoader) Load(materialDefinition MaterialDef) *Material {
	log.Println("Loading Material Definition", materialDefinition.Name)

	material := Material{
		Name: materialDefinition.Name,
		Shader: self.shaderLoader.Load(materialDefinition.Shaders),
	}

	if materialDefinition.Texture != "" {
		material.Texture = self.textureLoader.Load(materialDefinition.Texture)
	}

	self.loadedMaterials[material.Name] = &material

	return &material
}

// Get returns the Material with the given name. If no material has been loaded
// with the requested name, then returns the Default material
func (self *MaterialLoader) Get(materialName string) *Material {
	material, ok := self.loadedMaterials[materialName]

	if !ok {
		log.Println("Unable to find material", materialName, "defaulting to the default")
		material = self.loadedMaterials[defaultMaterialName]
	}

	return material
}
