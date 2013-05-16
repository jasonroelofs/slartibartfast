package main

import (
	"components"
	"core"
	"math3d"
	"render"
)

type TexturedCube struct {
	game    *Game
	theCube *core.Entity
}

func NewTexturedCube(game *Game) *TexturedCube {
	return &TexturedCube{
		game: game,
	}
}

func (self *TexturedCube) Setup() {
	self.theCube = core.NewEntity()
	self.theCube.AddComponent(new(components.Visual))

	self.game.graphicalBehavior.LoadMaterial(render.MaterialDef{
		Name:    "uvMap",
		Texture: "uvtemplate.tga",
		Shaders: "1texture_unlit",
	})

	transform := components.GetTransform(self.theCube)
	transform.Scale = math3d.Vector{5, 5, 5}

	visual := components.GetVisual(self.theCube)
	visual.MaterialName = "uvMap"

	self.game.RegisterEntity(self.theCube)
}

func (self *TexturedCube) Tick(deltaT float32) {
	// Do nothing
}
