package main

import (
	"behaviors"
	"configs"
	"core"
	"events"
	"input"
	"math3d"
	"platform"
	"render"
)

type Game struct {
	// Global management
	config          *configs.Config
	entityDB        *core.EntityDB
	inputDispatcher *input.InputDispatcher

	// Platform hooks
	renderer render.Renderer
	window   core.Window

	// The World
	camera *core.Camera

	// Behaviors
	graphicalBehavior *behaviors.Graphical
	inputBehavior     *behaviors.Input
	transformBehavior *behaviors.Transform

	// The Game
	player       *Player
	currentLevel *Level
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (self *Game) Run() {
	self.initializePlatform()
	self.initializeEngine()
	self.setupGame()
	self.mainLoop()
}

func (self *Game) initializePlatform() {
	self.window = platform.NewOpenGLWindow(self.config)
	self.renderer = platform.NewOpenGLRenderer()

	self.window.Open()
}

func (self *Game) initializeEngine() {
	self.entityDB = core.NewEntityDB()

	keyboard := input.NewKeyboard(self.window)
	mouse := input.NewMouse(self.window)
	self.inputDispatcher = input.NewInputDispatcher(self.config, keyboard, mouse)

	self.graphicalBehavior = behaviors.NewGraphical(self.renderer, self.entityDB)
	self.inputBehavior = behaviors.NewInput(self.inputDispatcher, self.entityDB)
	self.transformBehavior = behaviors.NewTransform(self.entityDB)

	self.loadBaseResources()
}

func (self *Game) loadBaseResources() {
	self.graphicalBehavior.LoadMaterial(render.MaterialDef{
		Name:    "only_color",
		Shaders: "color_unlit",
	})
}

func (self *Game) setupGame() {
	self.camera = core.NewCamera()
	self.camera.Perspective(60.0, self.window.AspectRatio(), 0.1, 100.0)
	self.camera.SetPosition(math3d.Vector{0, 10, 0})
	self.camera.LookAt(math3d.Vector{0, 0, -0.1})
	self.entityDB.RegisterEntity(self.camera.Entity)

	self.player = NewPlayer()
	self.entityDB.RegisterEntity(self.player.GetEntity())

	self.currentLevel = NewLevel()
	self.entityDB.RegisterEntity(self.currentLevel.Generate())
}

func (self *Game) mainLoop() {
	running := true
	self.inputDispatcher.On(events.Quit, func(e events.Event) {
		running = false
	})

	for running && self.window.IsOpen() {
		self.Tick(self.window.TimeSinceLast())
		self.window.SwapBuffers()
	}
}

func (self *Game) Tick(deltaT float32) {
	self.inputBehavior.Update(deltaT)
	self.transformBehavior.Update(deltaT)

	self.graphicalBehavior.Update(self.camera, deltaT)
}

func (self *Game) Shutdown() {
	self.window.Close()
}
