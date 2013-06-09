package factories

import (
	"components"
	"core"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_NewSkyBox(t *testing.T) {
	camera := core.NewCamera()
	skybox := NewSkyBox("testMap", camera)

	visual := components.GetVisual(skybox)
	transform := components.GetTransform(skybox)

	assert.Equal(t, "Skybox testMap", skybox.Name)
	assert.Equal(t, "testMap", visual.MaterialName)
	assert.Equal(t, camera.Entity, transform.UsingPositionOf)
}
