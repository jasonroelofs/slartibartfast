package main

import (
	"behaviors"
	"components"
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
	config    *configs.Config
	entityDB  *core.EntityDB
	behaviors []behaviors.Behavior
	renderer  render.Renderer
	window    core.Window

	boxen     []core.Entity
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
	graphical := behaviors.NewGraphical(self.renderer, self.entityDB)

	self.behaviors = append(self.behaviors, graphical)
}

func (self *Game) initializeScene() {
	var box *core.Entity

	positions := [10]math3d.Vector{
		math3d.Vector{0, 0, 0},
		math3d.Vector{4, 0, 0},
		math3d.Vector{4, 4, 0},
		math3d.Vector{4, 4, 4},
		math3d.Vector{8, 0, 0},
		math3d.Vector{8, 8, 0},
		math3d.Vector{0, 8, 0},
		math3d.Vector{4, 0, 4},
		math3d.Vector{0, 4, 4},
		math3d.Vector{0, 0, 4},
	}

	for i := 0; i < 10; i++ {
		box = core.NewEntityAt(positions[i])
		box.AddComponent(new(components.Visual))

		self.boxen = append(self.boxen, *box)

		self.RegisterEntity(box)
	}
}

func (self *Game) RegisterEntity(entity *core.Entity) {
	self.entityDB.RegisterEntity(entity)
}

// Figure out if I want float64 everywhere or not
func (self *Game) Tick(timeSinceLast float64) {
	deltaT := float32(timeSinceLast)
	box := self.boxen[0]

	// Transitioning!
	currPos := components.GetTransform(&box).Position
	t := components.GetTransform(&box)
	t.Position.X = currPos.X + deltaT
	t.Position.Y = currPos.Y + 2 * deltaT
	t.Position.Z = currPos.Z + 3 * deltaT

	if t.Position.X > 10 {
		t.Position.X = 0
	}

	if t.Position.Y > 10 {
		t.Position.Y = 0
	}

	if t.Position.Z > 10 {
		t.Position.Z = 0
	}

	// Scaling!
	box2 := self.boxen[1]
	t = components.GetTransform(&box2)
	t.Scale.X = t.Scale.X + deltaT
	t.Scale.Y = t.Scale.Y + 2 * deltaT
	t.Scale.Z = t.Scale.Z + 4 * deltaT

	if t.Scale.X > 3 {
		t.Scale.X = 1
	}

	if t.Scale.Y > 3 {
		t.Scale.Y = 1
	}

	if t.Scale.Z > 3 {
		t.Scale.Z = 1
	}

	// Rotating!
	box3 := self.boxen[2]
	t = components.GetTransform(&box3)
	t.Rotation = t.Rotation.RotateX(45.0 * deltaT)

	box4 := self.boxen[3]
	t = components.GetTransform(&box4)
	t.Rotation = t.Rotation.RotateY(90.0 * deltaT)

	box5 := self.boxen[4]
	t = components.GetTransform(&box5)
	t.Rotation = t.Rotation.RotateZ(135.0 * deltaT)

	box6 := self.boxen[5]
	t = components.GetTransform(&box6)
	t.Rotation = t.Rotation.RotateZ(45.0 * deltaT).RotateX(45.0 * deltaT)

	box7 := self.boxen[6]
	t = components.GetTransform(&box7)
	t.Rotation = t.Rotation.RotateY(45.0 * deltaT).RotateX(45.0 * deltaT)

	box8 := self.boxen[7]
	t = components.GetTransform(&box8)
	t.Rotation = t.Rotation.RotateZ(45.0 * deltaT).RotateY(45.0 * deltaT)

	box9 := self.boxen[8]
	t = components.GetTransform(&box9)
	t.Rotation = t.Rotation.RotateZ(45.0 * deltaT).RotateY(45.0 * deltaT).RotateX(45.0 * deltaT)

	for _, behavior := range self.behaviors {
		behavior.Update(timeSinceLast)
	}
}

func (self *Game) Shutdown() {
	self.window.Close()
}
