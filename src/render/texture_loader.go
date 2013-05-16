package render

import (
	"freeimage"
	"log"
)

type texturesMap map[string]*Texture

type TextureLoader struct {
	loadedTextures texturesMap
}

func NewTextureLoader() *TextureLoader {
	loader := new(TextureLoader)
	loader.loadedTextures = make(texturesMap)
	return loader
}

func (self *TextureLoader) Load(textureName string) *Texture {
	log.Println("Loading texture", textureName)

	texture := Texture{
		Image: freeimage.NewImage(Defaults.LoadPath + "/textures/" + textureName),
	}

	self.loadedTextures[textureName] = &texture

	return &texture
}
