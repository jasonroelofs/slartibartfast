package main

import (
	"components"
	"core"
	"math3d"
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
	self.theCube.Name = "Textured Cube"
	self.theCube.AddComponent(new(components.Visual))

	transform := components.GetTransform(self.theCube)
	transform.Position = math3d.Vector{0, 0, 0}
	transform.Scale = math3d.Vector{2, 2, 2}
	transform.Speed = math3d.Vector{5, 5, 5}

	self.game.RegisterEntity(self.theCube)
}

func (self *TexturedCube) Tick(deltaT float32) {
}
