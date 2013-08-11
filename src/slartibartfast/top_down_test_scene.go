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
			if y > 5 && (x > 3 && x < 47) && (z > 3 && z < 47) {
				return -1
			} else {
				return 1
			}
		},
	}

	volumeMesh := volume.MarchingCubes(self.levelVolume, math3d.Vector{50, 10, 50}, 0.5)
	volumeMesh.Name = "Level Mesh"

	self.levelEntity = core.NewEntity()
	self.levelEntity.Name = "Level Geometry"
	self.levelEntity.AddComponent(&components.Visual{
		Mesh:         volumeMesh,
		MaterialName: "only_color",
	})

	self.game.RegisterEntity(self.levelEntity)

	// Get the camera facing downwards
	cameraTransform := components.GetTransform(self.game.Camera.Entity)
	cameraTransform.Position = math3d.Vector{25, 10, 25}
	cameraTransform.CurrentPitch = 90

	// Replace input management to keep Camera in check
	self.game.Camera.RemoveComponent(components.INPUT)
	self.game.Camera.AddComponent(&components.Input{
		Mapping: FixedCameraMapping,
	})
}

func (self *TopDownTestScene) Tick(deltaT float32) {
}
