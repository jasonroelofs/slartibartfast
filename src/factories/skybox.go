package factories

import (
	"components"
	"core"
)

// NewSkyBox returns an Entity that is set up as a SkyBox. The material given
// needs to be the name of a cube-map Material definition. The resulting entity
// will be position-locked to the passed in camera so that it never looks like
// it's moving.
func NewSkyBox(skyboxMaterial string, camera *core.Camera) *core.Entity {
	entity := core.NewEntity()
	entity.Name = "Skybox " + skyboxMaterial

	visual := new(components.Visual)
	visual.MaterialName = skyboxMaterial

	transform := components.GetTransform(entity)
	transform.UsingPositionOf = camera.Entity

	entity.AddComponent(visual)

	return entity
}
