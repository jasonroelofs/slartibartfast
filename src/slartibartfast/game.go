package main

import (
	"behaviors"
	"components"
	"configs"
	"core"
	"events"
	"input"
	"platform"
)

type Game struct {
	config    *configs.Config
	entityDB  *core.EntityDB
	behaviors []behaviors.Behavior
	renderer  core.Renderer
	window    core.Window
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

	for running && self.window.IsOpen() {
		self.Tick()
		self.window.SwapBuffers()
	}
}

func (self *Game) initializeBehaviors() {
	graphical := behaviors.NewGraphical(self.renderer, self.entityDB)

	self.behaviors = append(self.behaviors, graphical)
}

func (self *Game) initializeScene() {
	box := core.NewEntity()
	box.AddComponent(new(components.Visual))

	self.RegisterEntity(box)
}

func (self *Game) RegisterEntity(entity *core.Entity) {
	self.entityDB.RegisterEntity(entity)
}

func (self *Game) Tick() {
	for _, behavior := range self.behaviors {
		behavior.Update(0)
	}
}

func (self *Game) Shutdown() {
	self.window.Close()
}
