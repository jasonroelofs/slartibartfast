package main

import (
	"components"
	"core"
	"events"
	"factories"
	"fmt"
	"input"
	"log"
	"math3d"
	"volume"
)

type VolumeScene struct {
	game *Game

	// How big of a cube, in world units, that MC will sample
	// from the volume
	marchingCubeSize float32

	cubeVolume   volume.Volume
	volumeEntity *core.Entity
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

	self.game.InputDispatcher.OnKey(input.KeyJ, func(e events.Event) {
		if e.Pressed {
			self.marchingCubeSize -= 0.1

			if self.marchingCubeSize <= 0 {
				self.marchingCubeSize = 0
			}

			self.rebuildVolume()
		}
	})

	self.game.InputDispatcher.OnKey(input.KeyK, func(e events.Event) {
		if e.Pressed {
			self.marchingCubeSize += 0.1

			self.rebuildVolume()
		}
	})

	self.cubeVolume = &volume.FunctionVolume{
		// A cube inside a 3x3x3 volume
		func(x, y, z float32) float32 {
			if x > 2.5 && x < 8.5 && y > 2.5 && y < 8.5 && z > 2.5 && z < 8.5 {
				return 1
			} else {
				return 0
			}
		},

		// A sphere!... or not. Need more control than 0 / 1
//		func(x, y, z float32) float32 {
//			if x*x + y*y + z*z < 9 { // 3^2
//				return 1
//			} else {
//				return 0
//			}
//		},
	}

	self.volumeEntity = core.NewEntity()
	self.volumeEntity.Name = "cube volume"

	self.rebuildVolume()

	// Move the volume into view of the starting camera
	transform := components.GetTransform(self.volumeEntity)
	transform.Position = math3d.Vector{-5, -5, -5}

	self.game.RegisterEntity(self.volumeEntity)
}

func (self *VolumeScene) Tick(deltaT float32) {
}

func (self *VolumeScene) rebuildVolume() {
	log.Println("Marching cube size:", self.marchingCubeSize)

	volumeMesh := volume.MarchingCubes(
		self.cubeVolume, math3d.Vector{10, 10, 10}, self.marchingCubeSize)
	volumeMesh.Name = fmt.Sprintf("CubeVolumeMesh[%.1f]", self.marchingCubeSize)

	self.volumeEntity.RemoveComponent(components.VISUAL)
	self.volumeEntity.AddComponent(&components.Visual{
		Mesh:         volumeMesh,
		MaterialName: "only_color",
	})
}
