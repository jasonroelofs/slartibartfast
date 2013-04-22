package main

import (
	"fmt"
	"configs"
)

func main() {
	fmt.Println("Welcome to Slartibartfast!")

	config, err := configs.NewConfig("config/settings.json")
	if err != nil {
		panic(fmt.Sprintf("Unable to read settings.json: %v", err))
	}

	game := NewGame(&config)
	game.Run()
	game.Shutdown()
}
