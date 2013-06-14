package main

import (
	"components"
	"core"
	"factories"
	"math3d"
)

type TexturedCube struct {
	game    *Game
	theCube *core.Entity

	colorCube *core.Entity
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

	self.colorCube = core.NewEntity()
	self.colorCube.Name = "Colored Cube"
	self.colorCube.AddComponent(&components.Visual{
		MaterialName: "only_color",
		MeshName: "ColorIndexCube",
	})

	transform := components.GetTransform(self.theCube)
	transform.Position = math3d.Vector{0, 0, -10}
	transform.Scale = math3d.Vector{2, 2, 2}
	transform.Speed = math3d.Vector{5, 5, 5}

	transform = components.GetTransform(self.colorCube)
	transform.Position = math3d.Vector{0, 0, -5}

	//	transform.MoveRelativeToRotation = true
	//	self.theCube.AddComponent(&components.Input{
	//		Mapping: FPSMapping,
	//	})

	skybox := factories.SkyBox("stevecube", self.game.Camera)

	self.game.RegisterEntity(skybox)
	self.game.RegisterEntity(self.theCube)
	self.game.RegisterEntity(self.colorCube)
}

func (self *TexturedCube) Tick(deltaT float32) {
}
