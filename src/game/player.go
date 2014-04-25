package main

import (
	"components"
	"core"
	"events"
	"math3d"
)

type Player struct {
	entity *core.Entity
}

func NewPlayer() *Player {
	player := new(Player)
	player.entity = core.NewEntity()
	player.entity.Name = "The Player"
	player.initializeComponents()

	return player
}

func (self *Player) initializeComponents() {
	topDownInput := &components.Input{
		Mapping: components.InputEventMap{
			events.MoveForward:  self.moveForward,
			events.MoveBackward: self.moveBackward,
			events.MoveLeft:     self.moveLeft,
			events.MoveRight:    self.moveRight,
			events.MouseMove:    self.faceMouse,
		},
		Polling: []events.EventType{
			events.MoveForward,
			events.MoveBackward,
			events.MoveLeft,
			events.MoveRight,
		},
	}

	self.entity.AddComponent(&components.Visual{})
	self.entity.AddComponent(topDownInput)

	transform := components.GetTransform(self.entity)
	transform.Scale = math3d.Vector{0.25, 0.5, 0.25}
	transform.Speed = math3d.Vector{3, 3, 3}
	transform.MoveRelativeToRotation = false
}

//
// Input component callbacks
//
func (self *Player) moveForward(_ components.ComponentHolder, event events.Event) {
	components.GetTransform(self.entity).MovingForward(event.Pressed)
}

func (self *Player) moveBackward(_ components.ComponentHolder, event events.Event) {
	components.GetTransform(self.entity).MovingBackward(event.Pressed)
}

func (self *Player) moveLeft(_ components.ComponentHolder, event events.Event) {
	components.GetTransform(self.entity).MovingLeft(event.Pressed)
}

func (self *Player) moveRight(_ components.ComponentHolder, event events.Event) {
	components.GetTransform(self.entity).MovingRight(event.Pressed)
}

func (self *Player) faceMouse(_ components.ComponentHolder, event events.Event) {
	transform := components.GetTransform(self.entity)

	transform.CurrentYaw = math3d.RadToDeg(
		math3d.Atan2(float32(event.MouseYDiff), float32(event.MouseXDiff)),
	)*-1 + 90
}

func (self *Player) GetEntity() *core.Entity {
	return self.entity
}
