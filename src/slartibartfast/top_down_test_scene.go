package main

import (
	"components"
	"core"
	"factories"
//	"log"
	"math3d"
	"volume"
)

// This scene is a top-down, fixed camera 2d esque movement test
// The "top-down" is set so that X / Y are the directions in movement and Z is depth

type TopDownTestScene struct {
	game *Game

	levelVolume volume.Volume
	levelEntity *core.Entity
}

func NewTopDownTestScene(game *Game) *TopDownTestScene {
	return &TopDownTestScene{
		game: game,
	}
}

func (self *TopDownTestScene) Setup() {
	skybox := factories.SkyBox("stevecube", self.game.Camera)
	self.game.RegisterEntity(skybox)

	self.levelVolume = &volume.FunctionVolume{
		func(x, y, z float32) float32 {
			if z > 5 {
				return -1
			} else {
				return 1
			}
		},
	}

	volumeMesh := volume.MarchingCubes(self.levelVolume, math3d.Vector{50, 50, 10}, 0.5)
	volumeMesh.Name = "Level Mesh"

	self.levelEntity = core.NewEntity()
	self.levelEntity.Name = "Level Geometry"
	self.levelEntity.AddComponent(&components.Visual{
		Mesh:         volumeMesh,
		MaterialName: "only_color",
	})

	self.game.RegisterEntity(self.levelEntity)

	self.game.Camera.SetPosition(math3d.Vector{25, 25, 10})

	self.game.Camera.RemoveComponent(components.INPUT)
	self.game.Camera.AddComponent(&components.Input{
		Mapping: FixedCameraMapping,
	})
}

func (self *TopDownTestScene) Tick(deltaT float32) {
}
