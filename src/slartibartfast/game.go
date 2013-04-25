package main

import (
	"behaviors"
	"components"
	"configs"
	"core"
	"events"
	"input"
	"window"
)

type Game struct {
	config    *configs.Config
	entityDB  core.EntityDB
	behaviors []behaviors.Behavior
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (self *Game) Run() {
	window.Open(self.config)

	self.initializeBehaviors()

	running := true

	inputHandler := input.NewInput()
	inputHandler.On(events.QUIT, func(e events.Event) {
		running = false
	})

	self.initializeScene()

	for running && window.StillOpen() {
		self.Tick()
		window.Present()
	}
}

func (self *Game) initializeBehaviors() {
	self.behaviors = append(self.behaviors, behaviors.NewGraphical(&self.entityDB))
}

func (self *Game) initializeScene() {
	box := core.NewEntity()
	box.AddComponent(components.Visual{})

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
