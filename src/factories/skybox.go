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
	// This mesh has no Texture coords or color settings because it is processed
	// as a cube map, and as such the Vertex positions double as texture coords.
	SkyBoxMesh = &render.Mesh{
		Name: "SkyBoxMesh",
		VertexList: []float32{
			-1.0, 1.0, -1.0,
			1.0, 1.0, -1.0,
			-1.0, -1.0, -1.0,
			1.0, -1.0, -1.0,
			-1.0, 1.0, 1.0,
			1.0, 1.0, 1.0,
			-1.0, -1.0, 1.0,
			1.0, -1.0, 1.0,
		},
		IndexList: []int32{
			2, 0, 3,
			3, 1, 0,
			3, 1, 7,
			7, 5, 1,
			6, 4, 2,
			2, 0, 4,
			7, 5, 6,
			6, 4, 5,
			0, 4, 1,
			1, 5, 4,
			6, 2, 7,
			7, 3, 2,
		},
	}
}
