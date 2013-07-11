package main

import (
	"components"
	"core"
	"factories"
	"math3d"
	"volume"
)

type VolumeScene struct {
	game *Game

	// How big of a cube, in world units, that MC will sample
	// from the volume
	marchingCubeSize float32
}

func NewVolumeScene(game *Game) *VolumeScene {
	return &VolumeScene{
		game:             game,
		marchingCubeSize: 1,
	}
}

func (self *VolumeScene) Setup() {
	skybox := factories.SkyBox("stevecube", self.game.Camera)
	self.game.RegisterEntity(skybox)

	cubeVolume := &volume.FunctionVolume{
		// A cube inside a 3x3x3 volume
		func(x, y, z float32) float32 {
			if x > 0.5 && x < 2.5 && y > 0.5 && y < 2.5 && z > 0.5 && z < 2.5 {
				return 1
			} else {
				return 0
			}
		},
	}

	volumeMesh := volume.MarchingCubes(cubeVolume, math3d.Vector{3, 3, 3}, self.marchingCubeSize)

	volumeEntity := core.NewEntity()
	volumeEntity.Name = "cube volume"
	volumeEntity.AddComponent(&components.Visual{
		Mesh: volumeMesh,
		MaterialName: "only_color",
	})

	self.game.RegisterEntity(volumeEntity)
}

func (self *VolumeScene) Tick(deltaT float32) {
}
