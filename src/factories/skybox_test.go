package factories

import (
	"components"
	"core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SkyBox(t *testing.T) {
	camera := core.NewCamera()
	skybox := SkyBox("testMap", camera)

	visual := components.GetVisual(skybox)
	transform := components.GetTransform(skybox)

	assert.Equal(t, "Skybox testMap", skybox.Name)
	assert.Equal(t, "testMap", visual.MaterialName)
	assert.Equal(t, camera.Entity, transform.UsingPositionOf)
}
