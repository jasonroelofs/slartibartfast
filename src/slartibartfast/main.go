package main

import (
	"configs"
	"log"
)

func main() {
	log.Println("Welcome to Slartibartfast!")

	config, err := configs.NewConfig("config/settings.json")
	if err != nil {
		log.Panicf("Unable to read settings.json: %v", err)
	}

	game := NewGame(config)
	game.Run()
	defer game.Shutdown()
}
