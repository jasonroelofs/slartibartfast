package behaviors

import (
	"core"
	"github.com/go-gl/gl"
)

type Graphical struct {
}

func (self *Graphical) Initialize(entityDB *core.EntityDB) {
}

func (self *Graphical) Update(deltaT float64) {
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Set the camera manually for now
	// A few units away and looking at the origin

	// Render the Visible's vertex / index arrays for each entity
}
