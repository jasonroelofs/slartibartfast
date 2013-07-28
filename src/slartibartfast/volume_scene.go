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
				self.marchingCubeSize = 0.1
			} else {
				self.rebuildVolume()
			}
		}
	})

	self.game.InputDispatcher.OnKey(input.KeyK, func(e events.Event) {
		if e.Pressed {
			self.marchingCubeSize += 0.1

			self.rebuildVolume()
		}
	})

	self.cubeVolume = &volume.FunctionVolume{
		// A cube
//		func(x, y, z float32) float32 {
//			if x > 20.5 && x < 38.5 && y > 20.5 && y < 38.5 && z > 20.5 && z < 38.5 {
//				return 1
//			} else {
//				return -1
//			}
//		},

		// A sphere!
		func(x, y, z float32) float32 {
			// Translate to treat middle of the volume as 0,0,0
			// Basically I want 0,0,0 of the volume to act like -25, -25, -25, so that
			// the center point of the sphere is at 25, 25, 25 in the volume,
			// then check the equation against the radius of the sphere.
			tX := x - 25
			tY := y - 25
			tZ := z - 25

			return -(tX*tX + tY*tY + tZ*tZ - 20)
		},
	}

	self.volumeEntity = core.NewEntity()
	self.volumeEntity.Name = "cube volume"

	self.rebuildVolume()

	// Move the volume into view of the starting camera
	transform := components.GetTransform(self.volumeEntity)
	transform.Position = math3d.Vector{-25, -25, -40}

	self.game.RegisterEntity(self.volumeEntity)
}

func (self *VolumeScene) Tick(deltaT float32) {
}

func (self *VolumeScene) rebuildVolume() {
	log.Println("Marching cube size:", self.marchingCubeSize)

	volumeMesh := volume.MarchingCubes(
		self.cubeVolume, math3d.Vector{50, 50, 50}, self.marchingCubeSize)
	volumeMesh.Name = fmt.Sprintf("CubeVolumeMesh[%.2f]", self.marchingCubeSize)

	self.volumeEntity.RemoveComponent(components.VISUAL)
	self.volumeEntity.AddComponent(&components.Visual{
		Mesh:         volumeMesh,
		MaterialName: "only_color",
	})
}
