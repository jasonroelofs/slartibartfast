package main

import (
	"behaviors"
	"components"
	"configs"
	"core"
	"events"
	"factories"
	"input"
	"log"
	"math3d"
	"platform"
	"render"
	"time"
)

type Game struct {
	config   *configs.Config
	entityDB *core.EntityDB

	renderer        render.Renderer
	window          core.Window
	Camera          *core.Camera
	InputDispatcher *input.InputDispatcher

	inputBehavior     *behaviors.Input
	transformBehavior *behaviors.Transform
	graphicalBehavior *behaviors.Graphical

	currentScene Scene
}

type Scene interface {
	Setup()
	Tick(deltaT float32)
	BeforeRender()
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (self *Game) Run() {
	self.window = platform.NewOpenGLWindow(self.config)
	self.window.Open()

	self.entityDB = core.NewEntityDB()
	self.renderer = new(platform.OpenGLRenderer)
	self.InputDispatcher = input.NewInputDispatcher()

	self.initializeBehaviors()
	self.loadAllMaterials()
	self.loadAllMeshes()
	self.initializeScene()

	running := true
	self.InputDispatcher.On(events.Quit, func(e events.Event) {
		running = false
	})

	var frameCount int64 = 0
	go calcAndPrintFPS(&frameCount)

	for running && self.window.IsOpen() {
		self.Tick(self.window.TimeSinceLast())
		self.window.SwapBuffers()
		frameCount++
	}
}

func calcAndPrintFPS(frameCount *int64) {
	ticker := time.Tick(1 * time.Second)
	var lastFrameCount int64 = *frameCount
	fps := 0

	for {
		<-ticker
		newCount := *frameCount
		fps = int(newCount - lastFrameCount)
		lastFrameCount = newCount

		log.Println("FPS:", fps)
	}
}

func (self *Game) initializeBehaviors() {
	self.inputBehavior = behaviors.NewInput(self.InputDispatcher, self.entityDB)
	self.transformBehavior = behaviors.NewTransform(self.entityDB)
	self.graphicalBehavior = behaviors.NewGraphical(self.renderer, self.entityDB)
}

// TODO also read from data/ to pull in any mesh files there.
func (self *Game) loadAllMeshes() {
	self.graphicalBehavior.LoadMesh(factories.SkyBoxMesh)

	colorCubeMesh := &render.Mesh{
		Name: "ColorIndexCube",
		VertexList: []float32{
			-1.0, 1.0, -1.0,
			1.0, 1.0, -1.0,
			-1.0, -1.0, -1.0,
			1.0, -1.0, -1.0,
			-1.0, 1.0, 1.0,
			1.0, 1.0, 1.0,
			-1.0, -1.0, 1.0,
			1.0, -1.0, 1.0,
		},
		ColorList: []float32{
			0.0, 0.0, 0.0,
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
			1.0, 1.0, 0.0,
			0.0, 1.0, 1.0,
			1.0, 0.0, 1.0,
			1.0, 1.0, 1.0,
		},
		IndexList: []int32{
			2, 0, 3,
			3, 1, 0,
			3, 1, 7,
			7, 5, 1,
			6, 4, 2,
			2, 0, 4,
			7, 5, 6,
			6, 4, 5,
			0, 4, 1,
			1, 5, 4,
			6, 2, 7,
			7, 3, 2,
		},
	}

	self.graphicalBehavior.LoadMesh(colorCubeMesh)
}

// TODO Read from data/ to get material definitions and send them
// down into the loader. For now this is a hard-coded set of materials
func (self *Game) loadAllMaterials() {
	// The stevecube space skybox texture from Ogre!
	self.graphicalBehavior.LoadMaterial(render.MaterialDef{
		Name:      "stevecube",
		Texture:   "stevecube.jpg",
		Shaders:   "cubemap",
		IsCubeMap: true,
	})

	self.graphicalBehavior.LoadMaterial(render.MaterialDef{
		Name: "only_color",
		Shaders: "color_unlit",
	})

}

func (self *Game) initializeScene() {
	self.Camera = core.NewCamera()
	self.Camera.Perspective(60.0, self.window.AspectRatio(), 0.1, 100.0)
	self.Camera.SetPosition(math3d.Vector{0, 0, 0})
	self.Camera.LookAt(math3d.Vector{0, 0, -5})

	input := components.Input{
		Mapping: FPSMapping,
	}

	self.Camera.AddComponent(&input)

//	self.Camera.SetSpeed(math3d.Vector{5, 5, 5})

	// YUCK, must be a better way of doing this?
	// Will probably move camera creation into Graphical so it
	// can take care of situations like this.
	self.RegisterEntity(self.Camera.Entity)

//	self.currentScene = NewSpinningCubes(self)
// self.currentScene = NewTexturedCube(self)
//	self.currentScene = NewVolumeScene(self)
	self.currentScene = NewTopDownTestScene(self)

	self.currentScene.Setup()
}

func (self *Game) RegisterEntity(entity *core.Entity) {
	self.entityDB.RegisterEntity(entity)
}

func (self *Game) Tick(deltaT float32) {
	self.currentScene.Tick(deltaT)

	self.inputBehavior.Update(deltaT)
	self.transformBehavior.Update(deltaT)

	self.currentScene.BeforeRender()
	self.graphicalBehavior.Update(self.Camera, deltaT)
}

func (self *Game) Shutdown() {
	self.window.Close()
}
