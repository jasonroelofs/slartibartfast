package main

import (
	"core"
	"components"
	"github.com/stretchr/testify/assert"
	"math3d"
	"testing"
)

func Test_NewTopDownCamera(t *testing.T) {
	camera := core.NewCamera()
	tdCam := NewTopDownCamera(camera)

	assert.NotNil(t, tdCam)
	assert.Equal(t, camera, tdCam.camera)
}

func Test_TopDownCamera_TracksEntities(t *testing.T) {
	camera := NewTopDownCamera(core.NewCamera())
	entity := core.NewEntity()
	camera.TrackEntity(entity)

	assert.Equal(t, entity, camera.trackingEntity)
}

func Test_TopDownCamera_UpdatesPositionToMatchEntity(t *testing.T) {
	baseCam := core.NewCamera()
	camera := NewTopDownCamera(baseCam)
	camera.TrackEntity(core.NewEntityAt(math3d.Vector{10, 10, 10}))

	camera.UpdatePosition()

	assert.Equal(t, math3d.Vector{10, 10, 10}, baseCam.Position())
}

func Test_TopDownCamera_UpdatePositionCanKeepAHeightFromTheEntity(t *testing.T) {
	baseCam := core.NewCamera()
	camera := NewTopDownCamera(baseCam)
	camera.TrackEntity(core.NewEntityAt(math3d.Vector{10, 10, 10}))
	camera.SetTrackingHeight(10)

	camera.UpdatePosition()

	assert.Equal(t, math3d.Vector{10, 20, 10}, baseCam.Position())
}

func Test_TopDownCamera_CanPauseTracking(t *testing.T) {
	baseCam := core.NewCamera()
	camera := NewTopDownCamera(baseCam)
	camera.TrackEntity(core.NewEntityAt(math3d.Vector{10, 10, 10}))
	camera.PauseTracking()

	camera.UpdatePosition()
	assert.Equal(t, math3d.Vector{0, 0, 0}, baseCam.Position())
}

func Test_TopDownCamera_ResumesTrackingPositionOverTime(t *testing.T) {
	baseCam := core.NewCamera()
	camera := NewTopDownCamera(baseCam)
	camera.TrackEntity(core.NewEntityAt(math3d.Vector{10, 10, 10}))
	camera.SetTrackingHeight(10)

	camera.ResumeTracking()

	assert.NotNil(t, baseCam.Entity.GetComponent(components.ANIMATION))
}
