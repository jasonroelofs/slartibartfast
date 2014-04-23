package main

import (
	"components"
	"core"
	"math3d"
	"volume"
)

type Level struct {
	volume volume.Volume
	entity *core.Entity
}

func NewLevel() *Level {
	return new(Level)
}

func (self *Level) Generate() *core.Entity {
	// A square arena with walls
	self.volume = &volume.FunctionVolume{
		func(x, y, z float32) float32 {
			if y > 5 && (x > 3 && x < 47) && (z > 3 && z < 47) {
				return -1
			} else {
				return 1
			}
		},
	}

	volumeMesh := volume.MarchingCubes(self.volume, math3d.Vector{50, 10, 50}, 0.5)
	volumeMesh.Name = "Level Mesh"

	self.entity = core.NewEntity()
	self.entity.Name = "Level Geometry"
	self.entity.AddComponent(&components.Visual{
		Mesh:         volumeMesh,
		MaterialName: "only_color",
	})

	transform := components.GetTransform(self.entity)
	transform.Position = math3d.Vector{-25, -10, -25}

	return self.entity
}
