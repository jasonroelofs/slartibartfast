package main

import (
	"components"
	"core"
	"math3d"
)

type SpinningCubes struct {
	game  *Game
	boxen []core.Entity
}

func NewSpinningCubes(game *Game) *SpinningCubes {
	scene := new(SpinningCubes)
	scene.game = game
	return scene
}

func (self *SpinningCubes) Setup() {
	var box *core.Entity

	positions := [10]math3d.Vector{
		math3d.Vector{0, 0, 0},
		math3d.Vector{4, 0, 0},
		math3d.Vector{4, 4, 0},
		math3d.Vector{4, 4, 4},
		math3d.Vector{-4, 0, 0},
		math3d.Vector{-4, -4, 0},
		math3d.Vector{0, -4, 0},
		math3d.Vector{-4, 0, -4},
		math3d.Vector{0, 4, 4},
		math3d.Vector{0, 0, 4},
	}

	for i := 0; i < 10; i++ {
		box = core.NewEntityAt(positions[i])
		box.AddComponent(new(components.Visual))

		self.boxen = append(self.boxen, *box)

		self.game.RegisterEntity(box)
	}
}

func (self *SpinningCubes) Tick(deltaT float32) {
	box := self.boxen[0]

	// Transitioning!
	currPos := components.GetTransform(&box).Position
	t := components.GetTransform(&box)
	t.Position.X = currPos.X + deltaT
	t.Position.Y = currPos.Y + 2*deltaT
	t.Position.Z = currPos.Z + 3*deltaT

	if t.Position.X > 10 {
		t.Position.X = 0
	}

	if t.Position.Y > 10 {
		t.Position.Y = 0
	}

	if t.Position.Z > 10 {
		t.Position.Z = 0
	}

	// Scaling!
	box2 := self.boxen[1]
	t = components.GetTransform(&box2)
	t.Scale.X = t.Scale.X + deltaT
	t.Scale.Y = t.Scale.Y + 2*deltaT
	t.Scale.Z = t.Scale.Z + 4*deltaT

	if t.Scale.X > 3 {
		t.Scale.X = 1
	}

	if t.Scale.Y > 3 {
		t.Scale.Y = 1
	}

	if t.Scale.Z > 3 {
		t.Scale.Z = 1
	}

	// Rotating!
	box3 := self.boxen[2]
	t = components.GetTransform(&box3)
	t.Rotation = t.Rotation.RotateX(45.0 * deltaT)

	box4 := self.boxen[3]
	t = components.GetTransform(&box4)
	t.Rotation = t.Rotation.RotateY(90.0 * deltaT)

	box5 := self.boxen[4]
	t = components.GetTransform(&box5)
	t.Rotation = t.Rotation.RotateZ(135.0 * deltaT)

	box6 := self.boxen[5]
	t = components.GetTransform(&box6)
	t.Rotation = t.Rotation.RotateZ(45.0 * deltaT).RotateX(45.0 * deltaT)

	box7 := self.boxen[6]
	t = components.GetTransform(&box7)
	t.Rotation = t.Rotation.RotateY(45.0 * deltaT).RotateX(45.0 * deltaT)

	box8 := self.boxen[7]
	t = components.GetTransform(&box8)
	t.Rotation = t.Rotation.RotateZ(45.0 * deltaT).RotateY(45.0 * deltaT)

	box9 := self.boxen[8]
	t = components.GetTransform(&box9)
	t.Rotation = t.Rotation.RotateZ(45.0 * deltaT).RotateY(45.0 * deltaT).RotateX(45.0 * deltaT)

}
