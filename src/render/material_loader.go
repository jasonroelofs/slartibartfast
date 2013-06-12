package render

import (
	"log"
	"path"
	"strings"
)

const (
	defaultMaterialName = ""
)

type materialMap map[string]*Material

type MaterialLoader struct {
	shaderLoader  *ShaderLoader
	textureLoader *TextureLoader

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
		Name:   materialDefinition.Name,
		Shader: self.shaderLoader.Load(materialDefinition.Shaders),
	}

	if materialDefinition.IsCubeMap {
		material.Texture = self.loadCubeMap(materialDefinition.Texture, &material)
	} else if materialDefinition.Texture != "" {
		material.Texture = self.textureLoader.Load(materialDefinition.Texture)
	}

	self.loadedMaterials[material.Name] = &material

	return &material
}

func (self *MaterialLoader) loadCubeMap(baseName string, material *Material) *Texture {
	log.Println("Loading CubeMap", baseName)

	fileExt := path.Ext(baseName)
	baseFilePath := strings.Replace(baseName, fileExt, "", -1)
	tl := self.textureLoader

	material.IsCubeMap = true

	material.CubeMap[CubeFace_Front] = tl.Load("cubemaps/" + baseFilePath + "_Front" + fileExt)
	material.CubeMap[CubeFace_Back] = tl.Load("cubemaps/" + baseFilePath + "_Back" + fileExt)

	material.CubeMap[CubeFace_Left] = tl.Load("cubemaps/" + baseFilePath + "_Left" + fileExt)
	material.CubeMap[CubeFace_Right] = tl.Load("cubemaps/" + baseFilePath + "_Right" + fileExt)

	material.CubeMap[CubeFace_Top] = tl.Load("cubemaps/" + baseFilePath + "_Top" + fileExt)
	material.CubeMap[CubeFace_Bottom] = tl.Load("cubemaps/" + baseFilePath + "_Bottom" + fileExt)

	return new(Texture)
}

// Get returns the Material with the given name. If no material has been loaded
// with the requested name, this returns the Default material.
func (self *MaterialLoader) Get(materialName string) *Material {
	material, ok := self.loadedMaterials[materialName]

	if !ok {
		log.Printf("Unable to find material %s, using default instead", materialName)
		material = self.loadedMaterials[defaultMaterialName]
		self.loadedMaterials[materialName] = material
	}

	return material
}
