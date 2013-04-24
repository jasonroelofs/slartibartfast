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
	config   *configs.Config
	entities []*core.Entity
	renderer behaviors.Graphical
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (game *Game) Run() {
	window.Open(game.config)

	running := true

	inputHandler := input.NewInput()
	inputHandler.On(events.QUIT, func(e events.Event) {
		running = false
	})

	game.initializeScene()

	for running && window.StillOpen() {
		game.Tick()
		window.Present()
	}
}

func (game *Game) initializeScene() {
	box := core.NewEntity()
	box.AddComponent(components.NewVisual())

	game.RegisterEntity(box)
}

func (game *Game) RegisterEntity(entity *core.Entity) {
	game.entities = append(game.entities, entity)
}

func (game *Game) Tick() {
	game.renderer.Tick(game.entities)
}

func (g *Game) Shutdown() {
	window.Close()
}
