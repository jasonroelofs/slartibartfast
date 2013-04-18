package main

import (
	"configs"
	"fmt"
	"time"
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

	// Game runs
	time.Sleep(5 * time.Second)
}
