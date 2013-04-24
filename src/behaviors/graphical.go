package behaviors

import (
	"core"
	"github.com/go-gl/gl"
)

type Graphical struct {
}

func (self *Graphical) Tick(entities []*core.Entity) {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Set the camera manually for now
	// A few units away and looking at the origin

	// Render the Visible's vertex / index arrays for each entity
}
