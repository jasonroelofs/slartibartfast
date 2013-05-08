package render

import (
	"math3d"
)

// A RenderOperation defines the data required for a single render call.
type RenderOperation struct {
	Mesh      *Mesh
	Material  *Material
	Transform math3d.Matrix
}
