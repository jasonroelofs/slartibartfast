package main

import (
	"configs"
	"events"
	"fmt"
	"github.com/go-gl/gl"
	"input"
	"window"
)

func main() {
	fmt.Println("Welcome to Slartibartfast!")

	config, err := configs.NewConfig("config/settings.json")
	if err != nil {
		panic(fmt.Sprintf("Unable to read settings.json: %v", err))
	}

	window.Open(&config)
	defer window.Close()

	running := true

	inputHandler := input.NewInput()
	inputHandler.On(events.QUIT, func(e events.Event) {
		running = false
	})

	for running && window.StillOpen() {
		// Move this
		gl.ClearColor(0, 0, 0, 0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		window.Present()
	}
}
