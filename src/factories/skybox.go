package factories

import (
	"components"
	"core"
	"math3d"
	"render"
)

var SkyBoxMesh *render.Mesh

// SkyBox returns an Entity that is set up as a SkyBox. The material given
// needs to be the name of a cube-map Material definition. The resulting entity
// will be position-locked to the passed in camera so that it never looks like
// it's moving.
func SkyBox(skyboxMaterial string, camera *core.Camera) *core.Entity {
	entity := core.NewEntity()
	entity.Name = "Skybox " + skyboxMaterial

	visual := new(components.Visual)
	visual.MeshName = "SkyBoxMesh"
	visual.MaterialName = skyboxMaterial

	transform := components.GetTransform(entity)
	transform.UsingPositionOf = camera.Entity
	transform.Scale = math3d.Vector{20, 20, 20}

	entity.AddComponent(visual)

	return entity
}

func init() {
	// Helpfully provided by
	// http://www.keithlantz.net/2011/10/rendering-a-skybox-using-a-cube-map-with-opengl-and-glsl/
	SkyBoxMesh = &render.Mesh{
		Name: "SkyBoxMesh",
		VertexList: []float32{
			-1.0, 1.0, 1.0,
			-1.0, -1.0, 1.0,
			1.0, -1.0, 1.0,
			1.0, 1.0, 1.0,
			-1.0, 1.0, -1.0,
			-1.0, -1.0, -1.0,
			1.0, -1.0, -1.0,
			1.0, 1.0, -1.0,
		},
		IndexList: []int32{
			0, 1, 2, 3,
			3, 2, 6, 7,
			7, 6, 5, 4,
			4, 5, 1, 0,
			0, 3, 7, 4,
			1, 2, 6, 5,
		},
	}
}
