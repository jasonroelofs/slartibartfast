package main

import (
	"components"
	"core"
)

type TopDownCamera struct {
	// The Entity being tracked by this camera
	trackingEntity *core.Entity

	// Distance above the Entity this camera stays
	trackingHeight float32

	// Link back to the actual camera object we move around
	camera *core.Camera

	currentlyTracking bool
}

func NewTopDownCamera(camera *core.Camera) *TopDownCamera {
	return &TopDownCamera{
		camera:            camera,
		currentlyTracking: true,
	}
}

// SetTrackingHeight sets the distance in world units of how far this camera
// will stay above the Entity being tracked (in the Y axis)
func (self *TopDownCamera) SetTrackingHeight(height float32) {
	self.trackingHeight = height
}

// TrackEntity takes the Entity this camera is supposed to keep it's eye on
func (self *TopDownCamera) TrackEntity(entity *core.Entity) {
	self.trackingEntity = entity
}

// PauseTracking turns off any position tracking of the current Entity
func (self *TopDownCamera) PauseTracking() {
	self.currentlyTracking = false
}

// ResumeTracking moves the camera back to the Entity being tracked and resumes
// camera position updates
func (self *TopDownCamera) ResumeTracking() {
	aboveEntity := components.GetTransform(self.trackingEntity).Position
	aboveEntity.Y += self.trackingHeight

	self.camera.AddComponent(
		components.NewPositionAnimation(
			aboveEntity, 0.5, func() { self.currentlyTracking = true },
		),
	)
}

// UpdatePosition calculates where this camera needs to be to stay tracking
// on the Entity, handling any potential changes of the Entity's position
func (self *TopDownCamera) UpdatePosition() {
	if self.currentlyTracking {
		entityPosition := components.GetTransform(self.trackingEntity).Position
		entityPosition.Y += self.trackingHeight

		self.camera.SetPosition(entityPosition)
	}
}
