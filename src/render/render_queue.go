package render

import (
	"math3d"
)

// A RenderQueue represents the full set of visual elements to be drawn for a given frame.
// One of these will be built per frame and sent to the renderer for rendering.
// It knows the camera, all entities, and any static geometry to be considered.
type RenderQueue struct {
	ProjectionMatrix math3d.Matrix
	ViewMatrix       math3d.Matrix

	renderOps []RenderOperation
}

func NewRenderQueue() *RenderQueue {
	return new(RenderQueue)
}

func (self *RenderQueue) Add(renderOp RenderOperation) {
	self.renderOps = append(self.renderOps, renderOp)
}

func (self *RenderQueue) RenderOperations() []RenderOperation {
	return self.renderOps
}
