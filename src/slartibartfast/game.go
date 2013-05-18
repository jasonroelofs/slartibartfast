package main

import (
	"behaviors"
	"configs"
	"core"
	"events"
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

	renderer     render.Renderer
	window       core.Window
	camera       *core.Camera
	currentScene Scene

	graphicalBehavior *behaviors.Graphical

	currentRotation float32
}

type Scene interface {
	Setup()
	Tick(deltaT float32)
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (self *Game) Run() {
	self.window = platform.NewOpenGLWindow(self.config)
	self.window.Open()

	self.entityDB = new(core.EntityDB)
	self.renderer = new(platform.OpenGLRenderer)

	self.initializeBehaviors()
	self.initializeScene()

	running := true

	inputHandler := input.NewInput()
	inputHandler.On(events.QUIT, func(e events.Event) {
		running = false
	})

	var frameCount int64 = 0
	go calcAndPrintFPS(&frameCount)

	for running && self.window.IsOpen() {
		self.Tick(self.window.TimeSinceLast())
		self.window.SwapBuffers()
		frameCount += 1
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
	self.graphicalBehavior = behaviors.NewGraphical(self.renderer, self.entityDB)
}

func (self *Game) initializeScene() {
	self.camera = core.NewCamera()
	// Needs to be window.Width() / window.Height()
	self.camera.Perspective(60.0, 4.0/3.0, 0.1, 100.0)
	self.camera.Position = math3d.Vector{20, 0, 20}
	self.camera.LookAt = math3d.Vector{0, 0, 0}

	self.currentScene = NewSpinningCubes(self)
//	self.currentScene = NewTexturedCube(self)

	self.currentScene.Setup()
}

func (self *Game) RegisterEntity(entity *core.Entity) {
	self.entityDB.RegisterEntity(entity)
}

func (self *Game) Tick(deltaT float32) {
	self.currentScene.Tick(deltaT)

	self.camera.Position = math3d.Vector{
		math3d.Cos(math3d.DegToRad(self.currentRotation)) * 20,
		self.camera.Position.Y,
		math3d.Sin(math3d.DegToRad(self.currentRotation)) * 20,
	}

	self.currentRotation += 30 * deltaT

	// Update all behaviors
	self.graphicalBehavior.Update(self.camera, deltaT)
}

func (self *Game) Shutdown() {
	self.window.Close()
}
