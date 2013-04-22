package main

import (
	"configs"
	"events"
	"github.com/go-gl/gl"
	"input"
	"window"
)

type Game struct {
	config *configs.Config
}

func NewGame(config *configs.Config) *Game {
	return &Game{config: config}
}

func (game *Game) Run() {
	window.Open(game.config)
	defer window.Close()

	running := true

	inputHandler := input.NewInput()
	inputHandler.On(events.QUIT, func(e events.Event) {
		running = false
	})

	for running && window.StillOpen() {
		game.Tick()
		window.Present()
	}
}

func (g *Game) Tick() {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (g *Game) Shutdown() {
}
