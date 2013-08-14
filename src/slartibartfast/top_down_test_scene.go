package main

import (
	"components"
	"core"
	"events"
	"factories"
	"input"
	"log"
	"math3d"
	"volume"
)

// This scene is a top-down, fixed camera 2d esque movement test
// The "top-down" is set so that X / Y are the directions in movement and Z is depth

type TopDownTestScene struct {
	game *Game

	levelVolume volume.Volume
	levelEntity *core.Entity
	playerCube  *core.Entity

	topDownCamera *TopDownCamera

	inputPlayer bool
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
	cameraTransform.Speed = math3d.Vector{3, 3, 3}

	// Our unit we'll control
	self.playerCube = core.NewEntityAt(math3d.Vector{25, 6, 25})
	self.playerCube.Name = "The Player"
	self.playerCube.AddComponent(&components.Visual{})

	playerTransform := components.GetTransform(self.playerCube)
	playerTransform.Scale = math3d.Vector{0.25, 0.5, 0.25}
	playerTransform.Speed = math3d.Vector{3, 3, 3}

	self.topDownCamera = NewTopDownCamera(self.game.Camera)
	self.topDownCamera.SetTrackingHeight(5)
	self.topDownCamera.TrackEntity(self.playerCube)

	self.game.RegisterEntity(self.playerCube)

	self.game.InputDispatcher.OnKey(input.KeySpace, func(event events.Event) { self.SwapInput(event) })

	// Start by controlling the player unit. Game defaults to controlling the camera
	self.SwapInput(events.Event{Pressed: true})
}

func (self *TopDownTestScene) SwapInput(event events.Event) {
	log.Println("Swap Input!")

	if !event.Pressed {
		return
	}

	if self.inputPlayer {
		self.playerCube.RemoveComponent(components.INPUT)
		self.game.Camera.AddComponent(&components.Input{
			Mapping: FixedCameraMapping,
		})

		self.inputPlayer = false
		self.topDownCamera.PauseTracking()

		log.Println("[Camera] Player now", self.playerCube)
	} else {
		self.game.Camera.RemoveComponent(components.INPUT)
		self.playerCube.AddComponent(&components.Input{
			Mapping: FixedYMapping,
		})

		self.inputPlayer = true
		self.topDownCamera.ResumeTracking()

		log.Println("[Player] Player now", self.playerCube)
	}
}

func (self *TopDownTestScene) Tick(deltaT float32) {
}

func (self *TopDownTestScene) BeforeRender() {
	self.topDownCamera.UpdatePosition()
}
