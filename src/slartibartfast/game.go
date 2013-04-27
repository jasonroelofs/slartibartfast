package main

import (
	"behaviors"
	"components"
	"configs"
	"core"
	"events"
	"input"
	"platform"
	"window"
)

type Game struct {
	config    *configs.Config
	entityDB  *core.EntityDB
	behaviors []behaviors.Behavior
	renderer  core.Renderer
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (self *Game) Run() {
	window.Open(self.config)

	self.initializeSystems()
	self.initializeBehaviors()
	self.initializeScene()

	running := true

	inputHandler := input.NewInput()
	inputHandler.On(events.QUIT, func(e events.Event) {
		running = false
	})

	for running && window.StillOpen() {
		self.Tick()
		window.Present()
	}
}

func (self *Game) initializeSystems() {
	self.entityDB = new(core.EntityDB)
	self.renderer = new(platform.OpenGLRenderer)
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
	window.Close()
}
