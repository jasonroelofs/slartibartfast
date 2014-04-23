package main

import (
	"components"
	"core"
)

type Player struct {
	entity *core.Entity
}

func NewPlayer() *Player {
	player := new(Player)
	player.entity = core.NewEntity()
	player.entity.Name = "The Player"
	player.entity.AddComponent(&components.Visual{})

	return player
}

func (self* Player) GetEntity() *core.Entity {
	return self.entity
}
